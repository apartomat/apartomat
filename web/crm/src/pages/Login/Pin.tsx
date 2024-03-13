import React, { useEffect, useState } from "react"
import { useNavigate } from "react-router-dom"

import { useConfirmLoginPin, ConfirmLoginPinMutation } from "./useConfirmLoginPin"

import { useAuthContext } from "shared/context/auth/context"
import { useToken } from "shared/browser/token"

import { Main, Heading, Box, Form, FormField, Button, MaskedInput, Text } from "grommet"

function Pin({ email, token, redirectTo = "/" }: { email: string; token: string; redirectTo?: string }) {
    const [confirmLogin, { data, loading }] = useConfirmLoginPin()
    const { reset, check } = useAuthContext()
    const [, saveToken] = useToken()

    const navigate = useNavigate()

    const [pin, setPin] = useState("")

    useEffect(() => {
        if (data?.confirmLoginPin.__typename === "LoginConfirmed") {
            saveToken(data?.confirmLoginPin?.token)
            reset()
            check()
            navigate(redirectTo)
        }
    }, [data, saveToken, check, redirectTo, history])

    function handleInputPin({ target: { value } }: React.ChangeEvent<HTMLInputElement>) {
        setPin(value)
    }

    function handleSubmit(event: React.FormEvent) {
        confirmLogin(token, pin)
        event.preventDefault()
    }

    return (
        <Main pad="large">
            <Box width="large">
                <Heading level={2}>Код отправлен</Heading>
                <Box>Для входа ввидете код, отправленый на {email}.</Box>
                <Box margin={{ vertical: "medium" }}>
                    <Form onSubmit={handleSubmit}>
                        <FormField label="Код" width="xsmall">
                            <MaskedInput
                                name="pin"
                                mask={[{ length: 6, placeholder: "• • • • • •", regexp: /^\d+$/ }]}
                                onChange={handleInputPin}
                                value={pin}
                                type="number"
                                pattern="[0-9]*"
                                inputMode="numeric"
                            />
                        </FormField>
                        <ConfirmPinError data={data} />
                        <Box direction="row" justify="between" margin={{ top: "medium" }}>
                            <Box>
                                <Button type="submit" primary label="Войти" disabled={loading} fill={false} />
                            </Box>
                        </Box>
                    </Form>
                </Box>
            </Box>
        </Main>
    )
}

function ConfirmPinError({ data }: { data: ConfirmLoginPinMutation | null | undefined }) {
    switch (data?.confirmLoginPin.__typename) {
        case "ExpiredToken":
            return (
                <Box pad={{ horizontal: "small" }}>
                    <Text color="status-error">Время истекло</Text>
                </Box>
            )
        case "InvalidToken":
            return (
                <Box pad={{ horizontal: "small" }}>
                    <Text color="status-error">Неверный код</Text>
                </Box>
            )
        default:
            return null
    }
}

export default Pin
