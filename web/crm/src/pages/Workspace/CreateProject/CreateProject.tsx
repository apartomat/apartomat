import React, { useEffect, useState } from "react"

import { useCreateProject, State as CreateProjectState } from "./useCreateProject"

import {
    Accordion,
    AccordionPanel,
    Box,
    Button,
    DateInput,
    Form,
    FormField,
    Heading,
    Layer,
    Text,
    TextInput,
} from "grommet"
import { FormClose } from "grommet-icons"

export default function CreateProject({
    workspaceId,
    setShow,
}: {
    workspaceId: string
    setShow: (show: boolean) => void
}) {
    const [name, setName] = useState("")
    const [create, , state] = useCreateProject()
    const [dates, setDates] = useState<string[]>([])

    const handleSubmit = async (event: React.FormEvent) => {
        event.preventDefault()
        let startAt = undefined,
            endAt = undefined

        if (dates.length > 1) {
            startAt = new Date(dates[0])
            endAt = new Date(dates[1])
        }

        create({ workspaceId, name, startAt, endAt })
    }

    const handleChangeName = (event: React.FormEvent<HTMLInputElement>) => {
        setName(event.currentTarget.value)
    }

    useEffect(() => {
        if (state.state === CreateProjectState.DONE) {
            setShow(false)
        }
    }, [state.state, setShow])

    const handleChangeDates = ({ value }: { value: string | string[] }) => {
        if (Array.isArray(value)) {
            setDates(value)
        }
    }

    return (
        <Layer>
            <Box pad="medium" gap="medium">
                <Box direction="row" justify="between" align="center">
                    <Heading level={2} margin="none">
                        Новый проект
                    </Heading>
                    <Button icon={<FormClose />} onClick={() => setShow(false)} />
                </Box>
                <Form onSubmit={handleSubmit} validate="submit">
                    {state.state === CreateProjectState.FAILED && <Text>{state.error.message}</Text>}
                    <FormField label="Название" htmlFor="input">
                        <TextInput onChange={handleChangeName} value={name} required />
                    </FormField>
                    <Accordion width="medium">
                        <AccordionPanel label="Даты">
                            <DateInput
                                inline
                                calendarProps={{
                                    daysOfWeek: false,
                                    firstDayOfWeek: 1, // Monday
                                    locale: "ru-RU",
                                }}
                                value={dates}
                                onChange={handleChangeDates}
                                width="medium"
                            />
                        </AccordionPanel>
                    </Accordion>
                    <Box direction="row" margin={{ top: "medium" }}>
                        <Button
                            type="submit"
                            primary
                            label="Создать"
                            disabled={state.state === CreateProjectState.CREATING}
                        />
                    </Box>
                </Form>
            </Box>
        </Layer>
    )
}
