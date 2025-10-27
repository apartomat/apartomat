import React, { useState } from "react"
import { Box, BoxExtendedProps, Button, Form, FormField, TextInput, LayerExtendedProps, CheckBox } from "grommet"
import { Modal, Header as ModalHeader } from "widgets/modal"
import { useAddSplitCoverToAlbum } from "./api"

export function AddSplitCover({
    albumId,
    onClickClose,
    onSplitCoverAdded,
    ...props
}: {
    albumId: string
    onClickClose?: () => void
    onSplitCoverAdded?: () => void
} & LayerExtendedProps) {
    const [addSplitCover, { loading, error, success }] = useAddSplitCoverToAlbum(albumId)

    const handleSubmit = async ({ value }: { value: any }) => {
        const year = value.year ? parseInt(value.year) : undefined
        const city = value.city || undefined
        const subtitle = value.subtitle || undefined

        const result = await addSplitCover({
            title: value.title.trim(),
            subtitle,
            imgFileId: value.imgFileId.trim(),
            withQr: value.withQr || false,
            city,
            year,
        })

        if (result.data?.addSplitCoverToAlbum?.__typename === "SplitCoverAdded" && onSplitCoverAdded) {
            onSplitCoverAdded()
        }
    }

    return (
        <Modal header="Добавить обложку" onClickClose={onClickClose} error={error} {...props}>
            <Form onSubmit={handleSubmit}>
                <Box gap="medium">
                    <FormField label="Заголовок" name="title" required>
                        <TextInput
                            name="title"
                            placeholder="Введите заголовок"
                        />
                    </FormField>

                    <FormField label="Подзаголовок" name="subtitle">
                        <TextInput
                            name="subtitle"
                            placeholder="Введите подзаголовок"
                        />
                    </FormField>

                    <FormField label="ID файла изображения" name="imgFileId" required>
                        <TextInput
                            name="imgFileId"
                            placeholder="Введите ID файла изображения"
                        />
                    </FormField>

                    <FormField label="Город" name="city">
                        <TextInput
                            name="city"
                            placeholder="Введите город"
                        />
                    </FormField>

                    <FormField label="Год" name="year">
                        <TextInput
                            name="year"
                            placeholder="Введите год"
                            type="number"
                        />
                    </FormField>

                    <FormField label="Добавить QR-код" name="withQr">
                        <CheckBox
                            name="withQr"
                            label="Добавить QR-код"
                        />
                    </FormField>

                    <Box direction="row" justify="between" margin={{ top: "medium" }}>
                        <Button
                            type="submit"
                            primary
                            busy={loading}
                            success={success}
                            label="Добавить обложку"
                        />
                        <Button label="Отмена" onClick={onClickClose} />
                    </Box>
                </Box>
            </Form>
        </Modal>
    )
}
