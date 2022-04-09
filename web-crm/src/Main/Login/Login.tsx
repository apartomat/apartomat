import React from "react";

import { Redirect } from "react-router-dom"
import LoginForm from "./LoginForm"
import CheckEmail from "./CheckEmail"
import Pin from "./Pin"

import useAuthContext, { UserContextStatus } from "../../common/context/auth/useAuthContext"

import { useLoginByEmail } from "./useLoginByEmail"

export function Login () {
    const [login, { data, loading, error }] = useLoginByEmail()
    const { user } = useAuthContext()

    switch (data?.loginByEmail.__typename) {
    case "LinkSentByEmail":
        return (
            <CheckEmail email={data?.loginByEmail.email} />
        )
    case "PinSentByEmail":
        return (
            <Pin email={data?.loginByEmail.email} token={data?.loginByEmail.token} />
        )
    case "ServerError":
        return (
            <LoginForm login={login} loading={loading} error={data?.loginByEmail} />
        )
    default:
        if (user.status === UserContextStatus.LOGGED) {
            return (
                <Redirect to={{ pathname: "/"}} />
            )
        }

        return (
            <LoginForm login={login} loading={loading} error={error} />
        )
    }
}

export default Login