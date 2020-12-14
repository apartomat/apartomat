import { createContext, useContext, useEffect, useState } from "react";
import { useProfileLazyQuery } from "../api/types.d";

export type UserContext =
    | UserCheckingContext
    | UserLoggedContext
    | UserUndefinedContext
    | UserUnauthorizedContext

export enum UserContextStatus {
    CHEKING = 'CHEKING',
    LOGGED = 'LOGGED',
    UNDEFINED = 'UNDEFINED',
    UNAUTHORIZED = 'UNAUTHORIZED'
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

type User = {
    email: string
    avatar: string
};

type UserUndefinedContext = {
    status: UserContextStatus.UNDEFINED
}

const userContextUndefined: UserUndefinedContext = {status: UserContextStatus.UNDEFINED};

export const authContext = createContext<{
    user: UserContext,
    check: () => void,
    reset: () => void,
    error: string | undefined
}>({ user: userContextUndefined, check: () => {}, reset: () => {}, error: undefined});

export function useAuthContext() {
    return useContext(authContext);
}

export function useAuthProvider() {
    const [user, setUser] = useState<UserContext>(userContextUndefined);
    const [loadProfile, { data,  error, loading }] = useProfileLazyQuery();

    function check() {
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
                    email: data?.profile.email,
                    avatar: data?.profile.gravatar?.url
                } as UserLoggedContext);
                break;
            case 'Forbidden':
                setUser({status: UserContextStatus.UNAUTHORIZED} as UserUnauthorizedContext);
                break;
        }
    }, [data, error, loading]);

    return { user, check, reset, error: error?.message };
}

export default useAuthContext;