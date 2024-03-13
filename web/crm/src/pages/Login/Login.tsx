import React, { useEffect } from "react"

import { useLoginByEmail } from "./useLoginByEmail"
import { useAuthContext, UserContextStatus } from "shared/context/auth/context"
import { useNavigate } from "react-router-dom"

import LoginForm from "./LoginForm"
import CheckEmail from "./CheckEmail"
import Pin from "./Pin"

export function Login() {
    const { user } = useAuthContext()

    const navigate = useNavigate()

    useEffect(() => {
        if (user.status === UserContextStatus.LOGGED || user.status === UserContextStatus.UNDEFINED) {
            navigate("/")
        }
    }, [user])

    const [login, { data, loading, error }] = useLoginByEmail()

    switch (data?.loginByEmail.__typename) {
        case "LinkSentByEmail":
            return <CheckEmail email={data?.loginByEmail.email} />
        case "PinSentByEmail":
            return <Pin email={data.loginByEmail.email} token={data.loginByEmail.token} />
        case "ServerError":
            return <LoginForm login={login} loading={loading} error={data?.loginByEmail} />
    }

    return <LoginForm login={login} loading={loading} error={error} />
}
