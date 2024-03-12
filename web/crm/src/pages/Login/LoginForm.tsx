import React, { useState } from "react"

import { Main, Heading, Form, FormField, TextInput, Box, Button, Text } from "grommet"

import { LoginByEmailFn } from "./useLoginByEmail"

function LoginForm({
    loading,
    login,
    error,
}: {
    loading: boolean
    login: LoginByEmailFn
    error?: { message: string }
}) {
    const [email, setEmail] = useState("")

    function handleInput(event: React.ChangeEvent<HTMLInputElement>) {
        setEmail(event.target.value)
    }

    function handleSubmit(event: React.FormEvent) {
        login(email)
        event.preventDefault()
    }

    return (
        <Main pad="large">
            <Box width="medium">
                <Heading level={2}>Войти</Heading>
                {error && (
                    <Box>
                        <Text>{error.message}</Text>
                    </Box>
                )}
                <Form onSubmit={handleSubmit}>
                    <FormField label="Электронная почта" value={email} onChange={handleInput}>
                        <TextInput name="email" type="email"></TextInput>
                    </FormField>
                    <Box direction="row" justify="between">
                        <Button type="submit" primary label="Войти" disabled={loading === true} />
                    </Box>
                </Form>
            </Box>
        </Main>
    )
}

export default LoginForm
