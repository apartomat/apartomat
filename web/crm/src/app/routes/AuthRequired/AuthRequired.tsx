import React, { useEffect } from "react"
import { Navigate, Outlet } from "react-router-dom"

import { Main, Box, Text, Heading, Paragraph } from "grommet"
import { Spinner } from "shared/ui/Spinner"

import { useAuthContext, UserContextStatus } from "context/auth"

export function AuthRequired() {
    const { user, check, error } = useAuthContext()

    useEffect(() => {
        if (user.status === UserContextStatus.UNDEFINED) {
            check()
        }
    }, [user, check])

    if (error !== undefined) {
        return (
            <Main pad="large">
                <Heading level={2}>Ошибка</Heading>
                <Paragraph>Не удалось получить профиль пользователя</Paragraph>
            </Main>
        )
    }

    switch (user.status) {
        case UserContextStatus.UNDEFINED:
        case UserContextStatus.CHEKING:
            return (
                <Main pad="large">
                    <Box direction="row" gap="small" align="center">
                        <Spinner message="Авторизация..." />
                        <Text>Авторизация...</Text>
                    </Box>
                </Main>
            )
        case UserContextStatus.SERVER_ERROR:
            return (
                <Main pad="large">
                    <Heading level={2}>Ошибка</Heading>
                    <Paragraph>Не удалось получить профиль пользователя</Paragraph>
                </Main>
            )
        case UserContextStatus.LOGGED:
            return <Outlet />
        default:
            return <Navigate to="/login" />
    }
}
