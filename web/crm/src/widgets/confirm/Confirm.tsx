import { Box, Button, Heading, Layer, LayerExtendedProps, Text } from "grommet"
import { Alert, FormClose } from "grommet-icons"

export function Confirm({
    header,
    text,
    confirmLabel,
    onClickClose,
    onCancel,
    onConfirm,
    confirmDisabled,
    confirmBusy,
    confirmSuccess,
    ...layerProps
}: {
    header: string
    text: string
    confirmLabel: string
    onClickClose?: () => void
    onCancel?: () => void
    onConfirm?: () => void
    confirmDisabled?: boolean
    confirmBusy?: boolean
    confirmSuccess?: boolean
} & LayerExtendedProps) {
    return (
        <Layer {...layerProps}>
            <Box pad="medium" gap="medium" width={{ min: "500px" }}>
                <Box direction="row" justify="between" align="center">
                    <Heading level={4} margin="none">
                        {header}
                    </Heading>
                    <Button plain icon={<FormClose />} size="small" onClick={onClickClose} />
                </Box>
                <Box
                    pad="small"
                    round="small"
                    direction="row"
                    gap="small"
                    align="center"
                    background={{ color: "status-warning", opacity: "weak" }}
                >
                    <Alert color="status-warning" size="medium" />
                    <Text>{text}</Text>
                </Box>
                <Box direction="row" justify="end" gap="small" margin={{ top: "small" }}>
                    <Button label="Отмена" onClick={onCancel} />
                    <Button
                        primary
                        label={confirmLabel}
                        busy={confirmBusy}
                        success={confirmSuccess}
                        disabled={confirmDisabled}
                        onClick={onConfirm}
                    />
                </Box>
            </Box>
        </Layer>
    )
}
