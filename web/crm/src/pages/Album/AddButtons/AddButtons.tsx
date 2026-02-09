import { useState, useRef } from "react"

import { Box, BoxExtendedProps, Button, Drop } from "grommet"
import { Add } from "grommet-icons"

export function AddButtons({
    onClickAddVisualizations,
    onClickUploadCover,
    onClickAddSplitCover,
    ...boxProps
}: {
    onClickAddVisualizations?: () => void
    onClickUploadCover?: () => void
    onClickAddSplitCover?: () => void
} & BoxExtendedProps) {
    const [open, setOpen] = useState(false)

    const targetRef = useRef<HTMLDivElement>(null)

    return (
        <Box {...boxProps}>
            <Box ref={targetRef} border={{ color: "background-front", size: "medium" }} round="large">
                <Button label="Добавить..." icon={<Add />} primary onClick={() => setOpen(true)} />
            </Box>

            {open && targetRef.current && (
                <Drop
                    elevation="none"
                    target={targetRef.current}
                    onClickOutside={() => setOpen(false)}
                    onEsc={() => setOpen(false)}
                    align={{ bottom: "bottom" }}
                    round="large"
                >
                    <Box gap="small" border={{ color: "background-front", size: "medium" }} direction="row">
                        <Button
                            primary
                            label="Загрузить обложку"
                            onClick={() => {
                                setOpen(false)
                                onClickUploadCover && onClickUploadCover()
                            }}
                        />
                        <Button
                            primary
                            label="Обложку"
                            color="accent-2"
                            onClick={() => {
                                setOpen(false)
                                onClickAddSplitCover && onClickAddSplitCover()
                            }}
                        />
                        <Button
                            primary
                            label="Визуализации"
                            color="accent-3"
                            onClick={() => {
                                setOpen(false)
                                onClickAddVisualizations && onClickAddVisualizations()
                            }}
                        />
                        <Button primary label="Ссылку на сайт" color="status-ok" />
                    </Box>
                </Drop>
            )}
        </Box>
    )
}
