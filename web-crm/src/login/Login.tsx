import React from "react";

import { Redirect } from "react-router-dom";
import Form from "./Form";
import CheckEmail from "./CheckEmail";

import { useLoginByEmail } from "./useLoginByEmail";
import useAuthContext, { UserContextStatus } from '../common/context/auth/useAuthContext';

export function Login () {
    const [login, { data, loading, error }] = useLoginByEmail();
    const { user } = useAuthContext();

    switch (data?.loginByEmail.__typename) {
    case "CheckEmail":
        const email = data?.loginByEmail.email;

        return (
            <CheckEmail email={email} />
        );
    case "ServerError":
        return (
            <Form login={login} loading={loading} error={data?.loginByEmail} />
        );

    default:
        if (user.status === UserContextStatus.LOGGED) {
            return (
                <Redirect to={{ pathname: "/"}} />
            );
        }

        return (
            <Form login={login} loading={loading} error={error} />
        );
    }
}

export default Login;