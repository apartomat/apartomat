import { ButtonExtendedProps } from "grommet"
import React from "react"
import { useEffect } from "react"

import useCreateAlbum from "./useCreateAlbum"

export default function CreateAlbumOnClick({
    projectId,
    children,
    onAlbumCreated,
}: {
    projectId: string
    children: React.ReactElement<ButtonExtendedProps>
    onAlbumCreated?: (id: string) => void
}): JSX.Element {
    const [create, { data }] = useCreateAlbum()

    const handleClick = () => {
        create(projectId, "Альбом")
    }

    useEffect(() => {
        switch (data?.createAlbum.__typename) {
            case "AlbumCreated":
                onAlbumCreated && onAlbumCreated(data?.createAlbum.album.id)
        }
    }, [data])

    return (
        <>
            {React.Children.map(children, (child) => {
                if (React.isValidElement(child)) {
                    return React.cloneElement(child, { onClick: handleClick })
                }

                return child
            })}
        </>
    )
}
