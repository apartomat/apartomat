import React, { useEffect, useState } from "react"

import { Box, BoxExtendedProps, RadioButtonGroup } from "grommet"
import { Document, Icon } from "grommet-icons"
import { RoundType } from "grommet/utils"

import { useChangeAlbumPageOrientation, PageOrientation as PageOrientationEnum } from "./useChangeAlbumPageOrientation"

export function PageOrientation({
    albumId,
    orientation,
    onChange,
    ...boxProps
}: {
    albumId: string,
    orientation: PageOrientationEnum,
    onChange?: (orientation: PageOrientationEnum) => void
} & BoxExtendedProps) {
    const options = [
        {value: PageOrientationEnum.Portrait, title: "portrait", icon: <Document/> },
        {value: PageOrientationEnum.Landscape, title: "landscape", icon: <Document transform="rotate(-90)"/> }
    ]

    const [ pageOrientation, setPageOrientation ] = useState<string>(orientation)

    const [ change, { data, error, loading } ] = useChangeAlbumPageOrientation(albumId)

    useEffect(() => {
        if (!data) {
            return
        }

        switch (data?.changeAlbumPageOrientation.__typename) {
            case "AlbumPageOrientationChanged":
                const { album } = data.changeAlbumPageOrientation
                setPageOrientation(album.settings.pageOrientation)
                break
            default:
                setPageOrientation(orientation)
        }
    }, [ data ])

    useEffect(() => {
        if (error) {
            setPageOrientation(orientation) 
        }
    }, [ error ])

    const handleOnChange = ({ target: { value }}: React.ChangeEvent<HTMLInputElement>) => {
        setPageOrientation(value)
        change(value as PageOrientationEnum)
    }

    return (
        <Box {...boxProps}>
            <RadioButtonGroup
                name="pageOrientation"
                direction="row"
                gap="none"
                options={options}
                value={pageOrientation}
                onChange={handleOnChange}
            >{({
                value,
                title,
                icon
            }: {
                value: string,
                title: string,
                icon: React.ReactNode
            }, { checked }: { checked: boolean, focus: boolean, hover: boolean }) => {
                return (
                    <Box
                        width="xxsmall"
                        height="xxsmall"
                        align="center"
                        justify="center"
                        background={checked ? "brand" : "background-contrast"}
                        round={roundBox(options, value)}
                        title={title}
                    >
                        {icon}
                    </Box>
                )
            }}</RadioButtonGroup>
        </Box>
    )
}

const roundBox = (options: { value: string }[], option: string): RoundType | undefined => {
    const opts = options.map(({ value }: { value: string }) => value)

    switch (opts.indexOf(option)) {
        case 0:
            return {corner: "left", size: "small"}
        case opts.length - 1:
            return {corner: "right", size: "small"}
        default:
            return undefined
    }
}
