import { useEffect, useState } from "react"
import { Box, Button, Form, FormField, Heading, Layer, LayerExtendedProps, MaskedInput, Select, Text } from "grommet"

import { FormClose } from "grommet-icons"

import { WorkspaceUserRoleDictionary, WorkspaceUserRoleDictionaryItem, WorkspaceUserRole } from "api/graphql"
import { useInviteUser } from "./useInviteUser"

export type Value = { email: string, role?: WorkspaceUserRole }

export default function Invite({
    workspaceId,
    roles,
    onInviteSent,
    onClickClose,
    ...layerProps
}: {
    workspaceId: string,
    roles: WorkspaceUserRoleDictionary,
    onInviteSent?: (to: string) => void,
    onClickClose?: () => void
} & LayerExtendedProps ) {
    const [ value, setValue ] = useState<Value>({ email: "", role: undefined })

    const [ errorMessage, setErrorMessage ] = useState<string | undefined>()

    const [ invite, { data, error } ] = useInviteUser(workspaceId)

    const handleSubmit = () => {
        const { email, role } = value

        if (email && role) {
            invite(email, role)
        }
    }

    useEffect(() => {
        switch (data?.inviteUser.__typename) {
            case "AlreadyInWorkspace":
                setErrorMessage("Пользователь уже приглашен")
                return
            case "Forbidden":
                setErrorMessage("Доступ запрещен")
                return
            case "NotFound":
            case "ServerError":
                setErrorMessage("Ошибка сервера")
                return
            case "InviteSent":
                onInviteSent && onInviteSent(data?.inviteUser.to)
                return
        }
    }, [ data ])

    useEffect(() => {
        if (error) {
            setErrorMessage("Ошибка сервера")
        }
    }, [ error ])

    return (
        <Layer {...layerProps}>
            <Box pad="medium" gap="medium" width="medium">
                <Box direction="row" justify="between" align="center">
                    <Heading level={2} margin="none">Пригласить</Heading>
                    <Button icon={ <FormClose/> } onClick={onClickClose}/>
                </Box>

                {errorMessage &&
                    <Box
                        pad="small"
                        round="small"
                        direction="row"
                        gap="small"
                        align="center"
                        background={{ color: "status-critical", opacity: "weak"}}
                    >
                        <Box border={{ color: "status-critical", size: "small"}} round="large">
                            <FormClose color="status-critical" size="medium"/>
                        </Box>
                        <Text weight="bold" size="medium">{errorMessage}</Text>
                    </Box>
                }

                <Box>
                    <Form
                        validate="change"
                        messages={{ required: "обязательное поле" }}
                        value={value}
                        onChange={(value) => setValue(value)}
                        onSubmit={handleSubmit}
                    >
                        <FormField
                            label="Электронная почта"
                            name="email"
                            required
                            validate={(val: string) => {
                                if (!val.match(/[\w\-_.]+@[\w\-_.]+\.[\w]+/)) {
                                    return { status: "error", message: "не соответствует формату" }
                                }
                            }}
                        >
                            <MaskedInput
                                name="email"
                                mask={[
                                    { regexp: /^[\w\-_.\+]+$/, placeholder: "example" },
                                    { fixed: '@' },
                                    { regexp: /^[\w\-_.]+$/, placeholder: "test" },
                                    { fixed: '.' },
                                    { regexp: /^[\w]+$/, placeholder: "org" },
                                ]}
                            />
                        </FormField>
                        <FormField label="Роль" name="role" required>
                            <Select
                                name="role"
                                options={roles.items.map(({ key, value }: WorkspaceUserRoleDictionaryItem) => {
                                    return {key, value}
                                } )}
                                valueKey={{ key: "key", reduce: true }}
                                labelKey="value"
                            />
                        </FormField>
                        <Box direction="row" margin={{ top: "medium" }}>
                            <Button type="submit" primary label="Отправить" />
                        </Box>
                    </Form>
                </Box>
            </Box>
        </Layer>
    )
}