import { Box, BoxExtendedProps, Button, Heading, Layer, LayerExtendedProps, Text } from "grommet"
import { FormClose } from "grommet-icons"
import React from "react"

export function Modal({
    children,
    onClickClose,
    header,
    error,
    ...props
}: {
    header?: string
    error?: string
    onClickClose?: () => void
} & LayerExtendedProps) {
    return (
        <Layer {...props}>
            <Box pad="medium" gap="small">
                {header && <Header onClickClose={onClickClose}>{header}</Header>}

                {error && <CriticalMessage>{error}</CriticalMessage>}

                {children}
            </Box>
        </Layer>
    )
}

export function Header({ children, onClickClose }: { children: React.ReactNode; onClickClose?: () => void }) {
    return (
        <Box direction="row" justify="between" align="center">
            <Heading level={3} margin="none">
                {children}
            </Heading>
            <Button icon={<FormClose />} onClick={onClickClose} />
        </Box>
    )
}

export function CriticalMessage({ children, ...props }: {} & BoxExtendedProps) {
    return (
        <Box
            pad="small"
            round="small"
            direction="row"
            gap="small"
            align="center"
            background={{ color: "status-critical", opacity: "weak" }}
            {...props}
        >
            <Box border={{ color: "status-critical", size: "small" }} round="large">
                <FormClose color="status-critical" size="medium" />
            </Box>
            <Text weight="bold" size="medium">
                {children}
            </Text>
        </Box>
    )
}
