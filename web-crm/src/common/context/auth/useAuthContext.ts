import { createContext, useContext, useEffect, useState } from "react";

import { useProfileLazyQuery } from "../../../api/types.d";

export type UserContext =
    | UserCheckingContext
    | UserLoggedContext
    | UserUndefinedContext
    | UserUnauthorizedContext
    | UserServerErrorContext

export enum UserContextStatus {
    CHEKING = 'CHEKING',
    LOGGED = 'LOGGED',
    UNDEFINED = 'UNDEFINED',
    UNAUTHORIZED = 'UNAUTHORIZED',
    SERVER_ERROR = 'SERVER_ERROR'
}

type UserCheckingContext = {
    status: UserContextStatus.CHEKING
}

type UserLoggedContext = {
    status: UserContextStatus.LOGGED
} & User;

type UserUnauthorizedContext = {
    status: UserContextStatus.UNAUTHORIZED
}

type UserServerErrorContext = {
    status: UserContextStatus.SERVER_ERROR
}

type User = {
    id: number
    email: string
    avatar: string
    defaultWorkspaceId: number
};

type UserUndefinedContext = {
    status: UserContextStatus.UNDEFINED
}

const userContextUndefined: UserUndefinedContext = {status: UserContextStatus.UNDEFINED};

const userEmpty: User = { id: 0, email: "", avatar: "", defaultWorkspaceId: 0};

export const authContext = createContext<{
    user: UserContext,
    concreteUser: User,
    check: () => void,
    reset: () => void,
    error: string | undefined
}>({ user: userContextUndefined, concreteUser: userEmpty, check: () => {}, reset: () => {}, error: undefined});

export function useAuthContext() {
    return useContext(authContext);
}

export function useAuthProvider() {
    const [user, setUser] = useState<UserContext>(userContextUndefined);
    const [loadProfile, { data,  error, loading }] = useProfileLazyQuery();
    const [concreteUser, setConcreteUser] = useState<User>(userEmpty);

    function check() {
        console.log("calling check......");
        if (user.status === UserContextStatus.UNDEFINED) {
            setUser({status: UserContextStatus.CHEKING} as UserCheckingContext);
            loadProfile();
        }
    }

    function reset() {
        setUser({status: UserContextStatus.UNDEFINED} as UserUndefinedContext);
    }

    useEffect(() => {
        switch (data?.profile.__typename) {
            case 'UserProfile':
                setUser({
                    status: UserContextStatus.LOGGED,
                    id: data?.profile.id,
                    email: data?.profile.email,
                    avatar: data?.profile.gravatar?.url,
                    defaultWorkspaceId: data?.profile.defaultWorkspace.id
                } as UserLoggedContext);

                setConcreteUser({
                    id: data?.profile.id,
                    email: data?.profile.email,
                    avatar: data?.profile.gravatar?.url,
                    defaultWorkspaceId: data?.profile.defaultWorkspace.id
                } as User);

                break;
            case 'Forbidden':
                setUser({status: UserContextStatus.UNAUTHORIZED} as UserUnauthorizedContext);
                break;
            case 'ServerError':
                setUser({status: UserContextStatus.SERVER_ERROR} as UserServerErrorContext);
                break;
        }
    }, [data, error, loading]);

    return { user, concreteUser, check, reset, error: error?.message };
}

export default useAuthContext;