import React, { useState, useEffect} from "react"

import { Layer, Box, Heading, Button, Calendar } from "grommet"
import { FormClose } from "grommet-icons"

import { useChangeProjectDates } from "./useProjectChangeDates"

export default function Change({
    projectId,
    startAt,
    endAt,
    onChange,
    onEsc,
    onClickOutside,
    onClickClose
}: {
    projectId: string,
    startAt?: string,
    endAt?: string,
    onChange?: (dates: { startAt?: string, endAt?: string }) => void,
    onEsc?: () => void,
    onClickOutside?: () => void,
    onClickClose?: () => void
}) {
    const [ dates, setDates ] = useState(startAt && endAt ? [[startAt, endAt]] : undefined)

    const [ change, { loading, data } ] = useChangeProjectDates()

    useEffect(() => {
        switch (data?.changeProjectDates.__typename) {
            case "ProjectDatesChanged":
                onChange && onChange({ startAt: dates && dates[0] && dates[0][0], endAt: dates && dates[0] && dates[0][1] })
        }
    }, [ data, dates, onChange ])


    /* eslint-disable  @typescript-eslint/no-explicit-any */
    const handleSelect = (value: any) => {
        setDates(value)
    }

    const handleSubmit = (event: React.FormEvent) => {
        change(projectId, { startAt: dates && dates[0] && dates[0][0], endAt: dates && dates[0] && dates[0][1] })
        event.preventDefault()
    }

    return (
        <Layer
            onClickOutside={onClickOutside}
            onEsc={onEsc}
        >
            <Box pad="medium" gap="medium">
                <Box direction="row" justify="between" align="center">
                    <Heading level={2} margin="none">Сроки проекта</Heading>
                    <Button icon={ <FormClose/> } onClick={onClickClose}/>
                </Box>
                <Box>
                    <Calendar
                        firstDayOfWeek={1}
                        locale="ru-RU"
                        range="array"
                        activeDate={undefined}
                        dates={dates}
                        onSelect={handleSelect}
                    />
                </Box>
                <Box direction="row" margin={{top: "medium"}}>
                    <Button type="submit" primary label="Сохранить" onClick={handleSubmit} disabled={loading}/>
                </Box>
            </Box>
        </Layer>
    )
}

