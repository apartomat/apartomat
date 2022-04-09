import React, { useEffect, useState } from "react";

import { useLocation, useHistory } from "react-router-dom";

import useConfirmLogin from "./useConfirmLogin";
import useAuthContext from "../common/context/auth/useAuthContext";
import useToken from "../common/context/auth/useToken";

export function Confirm({ redrectTo = "/"}) {
    const location = useLocation();
    const history = useHistory();
    const { check } = useAuthContext();
    const [, saveToken ] = useToken();

    const [confirmLogin, { data: confirmLoginResult, loading }] = useConfirmLogin();
    const [sent, setSent] = useState(false);

    useEffect(() => {
        const token = new URLSearchParams(location.search).get("token");
        if (token && confirmLoginResult === undefined && !sent && !loading) {
            confirmLogin(token);
            setSent(true);
        }
    }, [location, sent, setSent, confirmLogin, confirmLoginResult, loading]);

    useEffect(() => {
        if (confirmLoginResult?.confirmLoginLink.__typename === "LoginConfirmed") {
            saveToken(confirmLoginResult?.confirmLoginLink.token);
            check();
            history.push(redrectTo);
        }
    }, [confirmLoginResult, history, redrectTo, check, saveToken, loading])


    switch (confirmLoginResult?.confirmLoginLink.__typename) {
        case "InvalidToken":
            return (
                <div>
                    <h1>Invalid token</h1>
                    <p>Please <a href="/login">login</a> again</p>
                </div>
            );
        case "ServerError":
            return (
                <div>
                    <h1>Server error</h1>
                    <p>Please try again</p>
                </div>
            );
        default:
            return (
                <div>
                    <h1>Confirm login</h1>
                    <p>Please wait a moment...</p>
                </div>
            );
    }
}

export default Confirm;