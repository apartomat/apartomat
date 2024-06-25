import React, { ReactNode } from "react"

import { Box, BoxExtendedProps, Button, Header, Heading, Layer, Main, Text } from "grommet"

import { AnchorLink, Spinner } from "shared/ui"
import { UserAvatar } from "features/user-avatar"
import { Notifications } from "features/notification"

export function MainLayout({
    loading,
    error,
    notification,
    children,
    header,
    headerMenu,
}: {
    loading: boolean
    notification?: string
    header?: ReactNode
    headerMenu?: ReactNode
    error?: string
} & BoxExtendedProps) {
    return (
        <Main pad={{ vertical: "medium", horizontal: "large" }}>
            {loading && (
                <Layer position="top" margin="medium" animate={false} modal={false}>
                    <Box direction="row" gap="small">
                        <Spinner message="Загрузка..." />
                        <Text>Загрузка...</Text>
                    </Box>
                </Layer>
            )}

            <Notifications />

            {error && (
                <>
                    <Heading level={2}>Ошибка</Heading>
                    <Box>
                        <Text>{error}</Text>
                    </Box>
                </>
            )}

            {!error && (
                <>
                    <Header background="white" margin={{ vertical: "medium" }}>
                        <Box>
                            <Text size="xlarge" weight="bold" color="brand">
                                <AnchorLink to="/">apartomat</AnchorLink>
                            </Text>
                        </Box>
                        <Box>
                            <UserAvatar />
                        </Box>
                    </Header>

                    <Box direction="row" justify="between" margin={{ vertical: "medium" }}>
                        {header && (
                            <Heading level={2} margin="none">
                                {header}
                            </Heading>
                        )}
                        {headerMenu}
                    </Box>

                    <Box>{children}</Box>
                </>
            )}
        </Main>
    )
}
