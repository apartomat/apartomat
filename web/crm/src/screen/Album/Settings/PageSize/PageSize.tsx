import React, { useEffect, useState } from "react"

import { Box, BoxExtendedProps, RadioButtonGroup } from "grommet"
import { RoundType } from "grommet/utils"

import { useChangeAlbumPageSize, PageSize as PageSizeEnum } from "./useChangeAlbumPageSize"

const roundBox = (options: string[], option: string): RoundType | undefined => {
    switch (options.indexOf(option)) {
        case 0:
            return { corner: "left", size: "small" }
        case options.length - 1:
            return { corner: "right", size: "small" }
        default:
            return undefined
    }
}

export function PageSize({
    albumId,
    size,
    onAlbumPageSizeChanged,
    ...boxProps
}: {
    albumId: string
    size: PageSizeEnum
    onAlbumPageSizeChanged?: () => void
} & BoxExtendedProps) {
    const options = [PageSizeEnum.A3, PageSizeEnum.A4]

    const [pageSize, setPageSize] = useState<string>(size)

    const [change, { data, error }] = useChangeAlbumPageSize(albumId)

    useEffect(() => {
        if (!data) {
            return
        }

        switch (data?.changeAlbumPageSize.__typename) {
            case "AlbumPageSizeChanged": {
                const {
                    changeAlbumPageSize: { album },
                } = data
                setPageSize(album.settings.pageSize)
                onAlbumPageSizeChanged && onAlbumPageSizeChanged()
                break
            }
            default:
                setPageSize(size)
        }
    }, [data])

    useEffect(() => {
        if (error) {
            setPageSize(size)
        }
    }, [error])

    const handleOnChange = ({ target: { value } }: React.ChangeEvent<HTMLInputElement>) => {
        setPageSize(value)
        change(value as PageSizeEnum)
    }

    return (
        <Box {...boxProps}>
            <RadioButtonGroup
                name="pagesize"
                direction="row"
                gap="none"
                options={options}
                value={pageSize}
                onChange={handleOnChange}
            >
                {(option: string, { checked }: { checked: boolean; focus: boolean; hover: boolean }) => {
                    return (
                        <Box
                            width="xxsmall"
                            height="xxsmall"
                            align="center"
                            justify="center"
                            background={checked ? "brand" : "background-contrast"}
                            round={roundBox(options, option)}
                        >
                            {option}
                        </Box>
                    )
                }}
            </RadioButtonGroup>
        </Box>
    )
}
