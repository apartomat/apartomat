import { useEffect, useState } from "react"

import { useLocation, useNavigate } from "react-router-dom"

import useAcceptInvite from "./useAcceptInvite"
import { useAuthContext } from "shared/context/auth/context"
import { useToken } from "shared/browser/token"

export function AcceptInvite({ redirectTo = "/" }: { redirectTo?: string }) {
    const location = useLocation()
    const navigate = useNavigate()
    const { check } = useAuthContext()
    const [, saveToken] = useToken()

    const [confirmLogin, { data, loading }] = useAcceptInvite()

    const [sent, setSent] = useState(false)

    useEffect(() => {
        const token = new URLSearchParams(location.search).get("token")
        if (token && data === undefined && !sent && !loading) {
            confirmLogin(token)
            setSent(true)
        }
    }, [location, sent, setSent, confirmLogin, data, loading])

    useEffect(() => {
        switch (data?.acceptInvite.__typename) {
            case "InviteAccepted":
                saveToken(data?.acceptInvite.token)
                check()
                navigate(redirectTo)
                return
            case "AlreadyInWorkspace":
                navigate(redirectTo)
                return
        }
    }, [data, history, redirectTo, check, saveToken, loading])

    switch (data?.acceptInvite.__typename) {
        case "InvalidToken":
            return (
                <div>
                    <h1>Invalid token</h1>
                    <p>
                        Please <a href="/login">login</a> again
                    </p>
                </div>
            )
        case "ServerError":
            return (
                <div>
                    <h1>Server error</h1>
                    <p>Please try again</p>
                </div>
            )
        default:
            return (
                <div>
                    <h1>Confirm login</h1>
                    <p>Please wait a moment...</p>
                </div>
            )
    }
}
