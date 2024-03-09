import { useEffect } from "react"
import { useNavigate } from "react-router-dom"

import useToken from "context/auth/useToken"
import useAuthContext from "context/auth/useAuthContext"

function Logout() {
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

export default Logout
