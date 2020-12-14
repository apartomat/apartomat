import React, { useEffect, useState } from "react";

import { useLocation, useHistory } from "react-router-dom";

import useConfirmLogin from "./useConfirmLogin";
import useAuthContext from "../useAuthContext";
import useToken from "../useToken";

import Box from "@material-ui/core/Box";

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
        if (confirmLoginResult?.confirmLogin.__typename === "LoginConfirmed") {
            saveToken(confirmLoginResult?.confirmLogin.token);
            check();
            history.push(redrectTo);
        }
    }, [confirmLoginResult, history, redrectTo, check, saveToken, loading])


    switch (confirmLoginResult?.confirmLogin.__typename) {
        case "InvalidToken":
            return (
                <Box width={1/4}>
                    <h1>Invalid token</h1>
                    <p>Please <a href="/login">login</a> again</p>
                </Box>
            );
        case "ServerError":
            return (
                <Box width={1/4}>
                    <h1>Server error</h1>
                    <p>Please try again</p>
                </Box>
            );
        default:
            return (
                <Box width={1/4}>
                    <h1>Confirm login</h1>
                    <p>Please wait a moment...</p>
                </Box>
            );
    }
}

export default Confirm;