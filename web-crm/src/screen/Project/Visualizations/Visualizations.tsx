import React from "react"

import { Box, BoxExtendedProps, Image, Text } from "grommet"

import { ProjectFiles } from "../useProject"

export default function Visualizations({
    files,
    ...boxProps
}: {
    files: ProjectFiles
} & BoxExtendedProps) {
    switch (files.list.__typename) {
        case "ProjectFilesList":
            if (files.list.items.length === 0) {
                return null
            }

            return (
                <Box {...boxProps}>
                    <Box direction="row" gap="small" overflow={{"horizontal":"auto"}}>
                        {files.list.items.map(file =>
                            <Box
                                key={file.id}
                                height="small"
                                width="small"
                                margin={{bottom: "small"}}
                                flex={{"shrink":0}}
                                background="light-2"
                            >
                                <Image
                                    fit="cover"
                                    src={`${file.url}?w=192`}
                                    srcSet={`${file.url}?w=192 192w, ${file.url}?w=384 384w`}
                                />
                            </Box>
                        )}
                        {files.total.__typename === 'ProjectFilesTotal' && files.total.total > files.list.items.length
                            ? <Box key={0} height="small" width="small" margin={{bottom: "small"}} flex={{"shrink":0}} align="center" justify="center">
                                <Text>ещё {files.total.total - files.list.items.length}</Text>
                            </Box>
                            : null
                        }
                    </Box>
                </Box>
            )
        default:
            return null
    }
}