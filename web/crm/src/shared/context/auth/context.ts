import { createContext, useContext, useEffect, useMemo, useState } from "react"

import { useProfileLazyQuery } from "api/graphql"

export type UserContext =
    | UserCheckingContext
    | UserLoggedContext
    | UserUndefinedContext
    | UserUnauthorizedContext
    | UserServerErrorContext

export enum UserContextStatus {
    CHEKING = "CHEKING",
    LOGGED = "LOGGED",
    UNDEFINED = "UNDEFINED",
    UNAUTHORIZED = "UNAUTHORIZED",
    SERVER_ERROR = "SERVER_ERROR",
}

type UserCheckingContext = {
    status: UserContextStatus.CHEKING
}

type UserLoggedContext = {
    status: UserContextStatus.LOGGED
} & User

type UserUnauthorizedContext = {
    status: UserContextStatus.UNAUTHORIZED
}

type UserServerErrorContext = {
    status: UserContextStatus.SERVER_ERROR
}

type User = {
    id: string
    email: string
    avatar: string
    defaultWorkspaceId: string
}

type UserUndefinedContext = {
    status: UserContextStatus.UNDEFINED
}

const userContextUndefined: UserUndefinedContext = { status: UserContextStatus.UNDEFINED }

const userEmpty: User = { id: "", email: "", avatar: "", defaultWorkspaceId: "" }

export const authContext = createContext<{
    user: UserContext
    concreteUser: User
    check: () => void
    reset: () => void
    error: string | undefined
}>({ user: userContextUndefined, concreteUser: userEmpty, check: () => {}, reset: () => {}, error: undefined })

export function useAuthContext() {
    return useContext(authContext)
}

export function useAuthProvider() {
    const [user, setUser] = useState<UserContext>(userContextUndefined)

    const [loadProfile, { data, error, loading, refetch }] = useProfileLazyQuery()

    const [concreteUser, setConcreteUser] = useState<User>(userEmpty)

    function check() {
        if (user.status === UserContextStatus.UNDEFINED) {
            setUser({ status: UserContextStatus.CHEKING } as UserCheckingContext)

            if (data) {
                refetch()
            } else {
                loadProfile()
            }
        }
    }

    function reset() {
        setUser({ status: UserContextStatus.UNDEFINED } as UserUndefinedContext)
    }

    useEffect(() => {
        switch (data?.profile.__typename) {
            case "UserProfile":
                setUser({
                    status: UserContextStatus.LOGGED,
                    id: data?.profile.id,
                    email: data?.profile.email,
                    avatar: data?.profile.gravatar?.url,
                    defaultWorkspaceId: data?.profile.defaultWorkspace.id,
                } as UserLoggedContext)

                setConcreteUser({
                    id: data?.profile.id,
                    email: data?.profile.email,
                    avatar: data?.profile.gravatar?.url,
                    defaultWorkspaceId: data?.profile.defaultWorkspace.id,
                } as User)

                break
            case "Forbidden":
                setUser({ status: UserContextStatus.UNAUTHORIZED } as UserUnauthorizedContext)
                break
            case "ServerError":
                setUser({ status: UserContextStatus.SERVER_ERROR } as UserServerErrorContext)
                break
        }
    }, [data, error, loading])

    return useMemo(() => ({ user, concreteUser, check, reset, error: error?.message }), [user, concreteUser, error])
}
