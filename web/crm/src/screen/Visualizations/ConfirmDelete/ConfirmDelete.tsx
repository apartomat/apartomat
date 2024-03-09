import React from "react"

import { Box, Button, Heading, Layer, LayerExtendedProps, Text } from "grommet"
import { Alert, FormClose } from "grommet-icons"

export default function ConfirmDelete({
    count,
    disableButton,
    onClickClose,
    onClickDelete,
    ...layerProps
}: {
    count: number,
    disableButton: boolean,
    onClickClose?: () => void,
    onClickDelete?: () => void,
} & LayerExtendedProps) {
    return (
        <Layer
            {...layerProps}
        >
            <Box pad="medium" gap="medium" width={{min: "500px"}}>
                <Box direction="row" justify="between" align="center">
                    <Heading level={4} margin="none">Удалить {count} визуализации из проекта?</Heading>
                    <Button plain icon={ <FormClose/> } size="small" onClick={onClickClose}/>
                </Box>
                <Box
                    pad="small"
                    round="small"
                    direction="row"
                    gap="small"
                    align="center"
                    background={{ color: "status-warning", opacity: "weak"}}
                >
                    <Alert color="status-warning" size="medium"/>
                    <Text>Удаленные визуализации не возможно будет восстановить.</Text>
                </Box>
                <Box direction="row" justify="end" gap="small" margin={{ top: "small" }}>
                    <Button
                        label="Отмена"
                        onClick={onClickClose}
                    />
                    <Button
                        primary
                        label={disableButton ? "Удаление..." : "Удалить"}
                        disabled={disableButton}
                        onClick={onClickDelete}
                    />
                </Box>
            </Box>
        </Layer>
    )
}
