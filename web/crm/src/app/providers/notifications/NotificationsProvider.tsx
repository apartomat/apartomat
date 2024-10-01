import { ReactNode } from "react"

import { useNotificationsContextProvider, notificationsContext } from "shared/context/notifications/context"

export function NotificationsProvider({ children }: { children: ReactNode }) {
    const context = useNotificationsContextProvider()

    return <notificationsContext.Provider value={context}>{children}</notificationsContext.Provider>
}
