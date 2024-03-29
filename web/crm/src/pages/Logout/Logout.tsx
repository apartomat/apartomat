import { useEffect } from "react"
import { useNavigate } from "react-router-dom"

import { useToken } from "shared/browser/token"
import { useAuthContext } from "shared/context/auth/context"

export function Logout() {
    const [, , removeToken] = useToken()

    const { user, reset } = useAuthContext()

    const navigate = useNavigate()

    useEffect(() => {
        removeToken()
        reset()
    })

    useEffect(() => {
        navigate("/")
    }, [user])

    return <></>
}
