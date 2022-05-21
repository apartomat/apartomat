import React, { useState, useEffect } from "react"

import { Box, Button } from "grommet"

import ChangeDates from "../ChangeDates/ChangeDates"

export default function ProjectDates ({
    projectId,
    startAt,
    endAt,
    onProjectDatesChanged
}: {
    projectId: string,
    startAt?: string,
    endAt?: string,
    onProjectDatesChanged?: (dates: { startAt?: string, endAt?: string }) => void,
}) {
    const [ showChangeDates, setShowChangeDates ] = useState(false)

    const [ label, setLabel ] = useState(<>не определены</>)

    useEffect(() => {
        if (startAt && endAt) {
            return setLabel(
                <>
                    {new Date(startAt).toLocaleDateString("ru-RU")}&nbsp;&mdash;&nbsp;{new Date(endAt).toLocaleDateString("ru-RU")}
                </>
            )
        }

        if (startAt) {
            return setLabel(<>new Date(startAt).toLocaleDateString("ru-RU")</>)
        }

        return setLabel(<>не определены</>)
    }, [ startAt, endAt ])

    return (
        <>
            <Box direction="row">
                <Button
                    primary
                    color="light-2"
                    label={label}
                    onClick={() => setShowChangeDates(!showChangeDates)}
                />
            </Box>
            {showChangeDates &&
                <ChangeDates
                    projectId={projectId}
                    startAt={startAt}
                    endAt={endAt}
                    onEsc={() => setShowChangeDates(false) }
                    onClickOutside={() => setShowChangeDates(false) }
                    onClickClose={() => setShowChangeDates(false) }
                    onProjectDatesChanged={({ startAt, endAt }) => {
                        onProjectDatesChanged && onProjectDatesChanged({ startAt, endAt })
                        setShowChangeDates(false)
                    }}
                />
            }
        </>
    )
}