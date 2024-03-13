import { ReactNode } from "react"

import { authContext, useAuthProvider } from "shared/context/auth/context"

export function AuthProvider({ children }: { children: ReactNode }) {
    const auth = useAuthProvider()

    return <authContext.Provider value={auth}>{children}</authContext.Provider>
}
