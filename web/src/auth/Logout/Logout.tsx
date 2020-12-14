import React, { useEffect } from "react"
import { Redirect } from "react-router-dom";

import useToken from "../useToken";
import useAuthContext from "../useAuthContext";

function Logout() {
    const [,, removeToken] = useToken();
    const { reset } = useAuthContext();

    useEffect(() => {
        removeToken();
        reset();
    });

    return (
        <Redirect to="/"/>
    )
}

export default Logout;