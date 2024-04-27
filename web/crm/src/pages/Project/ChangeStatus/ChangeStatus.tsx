import React, { useState, useEffect, useMemo, useRef } from "react"

import { useChangeStatus } from "./useChangeProjectStatus"

import { ProjectStatus, ProjectStatusDictionary, ProjectStatusDictionaryItem } from "api/graphql"

import { Box, BoxExtendedProps, Button, Drop, Text } from "grommet"

export default function ChangeStatus({
    projectId,
    status,
    values,
    onProjectStatusChanged,
    onNotFound,
    onForbidden,
    onServerError,
    ...boxProps
}: {
    projectId: string
    status: ProjectStatus
    values?: ProjectStatusDictionary
    onProjectStatusChanged?: ({ status }: { status: ProjectStatus }) => void
    onNotFound?: () => void
    onForbidden?: () => void
    onServerError?: () => void
} & BoxExtendedProps) {
    const [showDropMenu, setShowDropMenu] = useState<boolean>(false)
    const [disabled, setDisabled] = useState(false)

    const [state, setState] = useState(status)

    const [changeStatus, { data, loading, error }] = useChangeStatus()

    const handleItemClick = (projectId: string, status: ProjectStatus) => {
        void changeStatus(projectId, status)
        setShowDropMenu(false)
        setState(status)
    }

    const rollback = (status: ProjectStatus) => {
        setTimeout(() => {
            setDisabled(loading)
            setState(status)
        }, 300)
    }

    useEffect(() => {
        if (loading) {
            setDisabled(loading)
        }

        switch (data?.changeProjectStatus.__typename) {
            case "ProjectStatusChanged": {
                const {
                    changeProjectStatus: {
                        project: { status },
                    },
                } = data
                setDisabled(loading)
                onProjectStatusChanged && onProjectStatusChanged({ status })

                break
            }
            case "NotFound":
                rollback(status)
                onNotFound && onNotFound()
                break
            case "Forbidden":
                rollback(status)
                onForbidden && onForbidden()
                break
            case "ServerError":
                rollback(status)
                onServerError && onServerError()
                break
            default:
                error ? rollback(status) : setDisabled(loading)
        }
    }, [data, error, loading])

    const label = useMemo(() => statusToLabel({ status: state, items: values?.items }), [state, values])

    const color = useMemo(() => statusColor(state), [state])

    const targetRef = useRef<HTMLDivElement>(null)

    return (
        <Box {...boxProps} justify="center">
            <Box ref={targetRef}>
                <Button
                    label={label}
                    color={color}
                    size="small"
                    onClick={() => setShowDropMenu(true)}
                    disabled={disabled}
                />
            </Box>

            {showDropMenu && targetRef.current && (
                <Drop
                    elevation="small"
                    round="small"
                    align={{ top: "bottom", left: "left" }}
                    margin={{ top: "xsmall" }}
                    target={targetRef.current}
                    onClickOutside={() => setShowDropMenu(false)}
                    onEsc={() => setShowDropMenu(false)}
                >
                    {values?.items.map((item) => {
                        return (
                            <Button key={item.key} plain hoverIndicator={{ color: "light-2" }}>
                                <Box pad="small" onClick={() => handleItemClick(projectId, item.key)}>
                                    <Text>{item.value}</Text>
                                </Box>
                            </Button>
                        )
                    })}
                </Drop>
            )}
        </Box>
    )
}

function statusToLabel({ status, items }: { status: ProjectStatus; items?: ProjectStatusDictionaryItem[] }): string {
    if (!items) {
        return ""
    }

    for (const item of items) {
        if (item.key === status) {
            return item.value
        }
    }

    return ""
}

function statusColor(status: ProjectStatus): string {
    switch (status) {
        case ProjectStatus.New:
            return "status-unknown"
        case ProjectStatus.InProgress:
            return "status-ok"
        case ProjectStatus.Done:
            return "status-warning"
        case ProjectStatus.Canceled:
            return "status-error"
    }
}
