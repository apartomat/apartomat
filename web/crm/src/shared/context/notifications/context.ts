import { createContext, useContext, useState } from "react"

export const notificationsContext = createContext<{
    notifications: NotificationMessage[]
    notify: (message: message) => void
    dismiss: (id: string) => void
}>({
    notifications: [],
    notify: (message: message) => {},
    dismiss: (id: string) => {},
})

type message = {
    message: string
    severity?: "unknown" | "ok" | "warning" | "error" | "critical"
    timeout?: number
    duration?: number
    callback?: () => void
}

export type NotificationMessage = {
    id: string
    message: string
    severity: "unknown" | "ok" | "warning" | "error" | "critical"
    timeout: number
    duration: number
}

export function useNotifications() {
    return useContext(notificationsContext)
}

export function useNotificationsContextProvider() {
    const [notifications, setNotifications] = useState<NotificationMessage[]>([])

    return {
        notifications,
        notify: ({
            message,
            severity = "ok",
            timeout = 250,
            duration = 1000,
            callback,
        }: message) => {
            const id = Math.random().toString(36).slice(2, 9) + new Date().getTime().toString(36)

            if (callback) {
                setTimeout(callback, timeout + duration)
            }

            setNotifications((prev) => [...prev, { id, message, severity, timeout, duration, callback } as NotificationMessage])
        },
        dismiss: (id: string) => {
            setNotifications((prev) => prev.filter((message) => message.id !== id))
        },
    }
}
