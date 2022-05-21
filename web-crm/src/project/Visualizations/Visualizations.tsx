import React, { Dispatch, SetStateAction } from "react"

import { Box, Heading, Image, Button, Text } from "grommet"

import { ProjectFiles } from "../useProject"

export default function Visualizations({ files, showUploadFiles }: { files: ProjectFiles, showUploadFiles: Dispatch<SetStateAction<boolean>> }) {
    const handleUploadFiles = () => {
        showUploadFiles(true)
    }

    switch (files.list.__typename) {
        case "ProjectFilesList":
            if (files.list.items.length === 0) {
                return null
            }

            return (
                <Box margin={{vertical: "medium"}}>
                    <Box direction="row" justify="between">
                        <Heading level={3}>Визуализации</Heading>
                        <Box justify="center">
                            <Button color="brand" label="Загрузить" onClick={handleUploadFiles} />
                        </Box>
                    </Box>
                    <Box direction="row" gap="small" overflow={{"horizontal":"auto"}} >
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
            return <div>n/a</div>
    }
}