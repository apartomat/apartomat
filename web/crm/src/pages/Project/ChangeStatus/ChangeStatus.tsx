import React, { useState, useEffect, useMemo, useRef } from "react"

import { useChangeStatus } from "./useChangeProjectStatus"

import { ProjectStatus, ProjectStatusDictionary, ProjectStatusDictionaryItem } from "api/graphql"

import { Box, BoxExtendedProps, Button, Drop, Text } from "grommet"

export default function ChangeStatus({
    projectId,
    status,
    values,
    onProjectStatusChanged,
    ...boxProps
}: {
    projectId: string
    status: ProjectStatus
    values?: ProjectStatusDictionary
    onProjectStatusChanged?: ({ status }: { status: ProjectStatus }) => void
} & BoxExtendedProps) {
    const [show, setShow] = useState<boolean>(false)

    const [state, setState] = useState(status)

    const [changeStatus, { data, loading, error }] = useChangeStatus()

    const handleItemClick = (projectId: string, status: ProjectStatus) => {
        changeStatus(projectId, status)
        setShow(false)
        setState(status)
    }

    useEffect(() => {
        switch (data?.changeProjectStatus.__typename) {
            case "ProjectStatusChanged": {
                const {
                    changeProjectStatus: {
                        project: { status },
                    },
                } = data
                onProjectStatusChanged && onProjectStatusChanged({ status })
            }
        }
    }, [data, onProjectStatusChanged])

    useEffect(() => {
        if (error) {
            setTimeout(() => setState(status), 300)
        }
    })

    const label = useMemo(() => statusToLabel({ status: state, items: values?.items }), [state, values])

    const color = useMemo(() => statusColor(state), [state])

    const targetRef = useRef<HTMLDivElement>(null)

    return (
        <Box {...boxProps} justify="center">
            <Box ref={targetRef}>
                <Button label={label} color={color} size="small" onClick={() => setShow(true)} disabled={loading} />
            </Box>

            {show && targetRef.current && (
                <Drop
                    elevation="small"
                    round="small"
                    align={{ top: "bottom", left: "left" }}
                    margin={{ top: "xsmall" }}
                    target={targetRef.current}
                    onClickOutside={() => setShow(false)}
                    onEsc={() => setShow(false)}
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
