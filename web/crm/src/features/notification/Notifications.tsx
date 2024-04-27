import { Box, Layer, Text } from "grommet"
import { StatusGood, StatusCritical, CircleAlert, StatusWarning, StatusInfo } from "grommet-icons"
import React, { useEffect, useState } from "react"
import { useNotifications, NotificationMessage } from "shared/context/notiifcations/context"

export function Notifications() {
    const { notifications, dismiss } = useNotifications()

    return (
        notifications.length > 0 && (
            <Layer position="top" modal={false} responsive={false} margin={{ vertical: "small", horizontal: "small" }}>
                {notifications.map((val: NotificationMessage, index: number) => (
                    <Message key={val.id} val={val} dismiss={() => dismiss(val.id)} />
                ))}
            </Layer>
        )
    )
}

function Message({ val, dismiss }: { val: NotificationMessage; dismiss: (id: string) => void }) {
    const [show, setShow] = useState(false)

    useEffect(() => {
        const timer = setTimeout(() => {
            setShow(true)

            setTimeout(() => {
                dismiss(val.id)
            }, val.duration || 1000)
        }, val.timeout || 1250)

        return () => clearTimeout(timer)
    }, [])

    return (
        show && (
            <Box
                align="center"
                direction="row"
                gap="xsmall"
                justify="between"
                elevation="small"
                background={`status-${val.severity}`}
                round="medium"
                pad={{ vertical: "xsmall", horizontal: "small" }}
                animation={"zoomIn"}
            >
                <Icon severity={val.severity} />
                <Text>{val.message}</Text>
            </Box>
        )
    )
}

function Icon({ severity }: { severity: "unknown" | "ok" | "warning" | "error" | "critical" }) {
    switch (severity) {
        case "critical":
            return <CircleAlert />
        case "error":
            return <StatusCritical />
        case "warning":
            return <StatusWarning />
        case "ok":
            return <StatusGood />
        default:
            return <StatusInfo />
    }
}
