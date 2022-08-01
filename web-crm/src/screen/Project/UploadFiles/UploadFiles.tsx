import React, { useState, useEffect, Dispatch, SetStateAction, ChangeEvent } from "react"

import { Layer, Form, FormField, Box, Heading, Button, FileInput, Image, Text  } from "grommet"
import { FormClose } from "grommet-icons"

import useUploadProjectFile, { ProjectFileType } from "./useUploadProjectFile"

export default function UploadFiles(
    { projectId, type, setShow, onUploadComplete }:
    {
        projectId: string,
        type: ProjectFileType,
        setShow: Dispatch<SetStateAction<boolean>>,
        onUploadComplete: ({message}: { message: string}) => void
    }
) {
    const [ files, setFiles ] = useState<FileList | null>(null)
    const [ upload, { loading, error, data, called }, state ] = useUploadProjectFile()
    
    useEffect(() => {
        if (state.state === "DONE") {
            console.log(state.file);
        }

        if (state.state === "FAILED") {
            if (state.error instanceof Error) {
                console.log("------------", state.error.message)
            } else {
                console.log(state.error.__typename, state.error.message)
            }
            
        }
    })

    const handleSubmit = (event: React.FormEvent) => {
        if (files && !loading) {
            upload({ projectId, file: files[0], type })
        }

        event.preventDefault()
    }

    const handleFileInputChange = (event: ChangeEvent<HTMLInputElement>) => {
        if (event.target.files) {
            setFiles(event.target.files)
        }
    }

    useEffect(() => {
        if (called && data && !error) {
            setShow(false)
            onUploadComplete({ message: files?.length === 1 ? "Файл загружен" : `Загружено файлов ${files?.length}`})
        }
    })

    return (
        <Layer
            onClickOutside={() => setShow(false)}
            onEsc={() => setShow(false)}
        >
            <Form validate="submit" onSubmit={handleSubmit}>
                <Box pad="medium" gap="medium" width="large">
                    <Box direction="row"justify="between"align="center">
                        <Heading level={3} margin="none">Загрузить файлы</Heading>
                        <Button icon={ <FormClose/> } onClick={() => setShow(false)}/>
                    </Box>
                    <FormField name="files" htmlFor="files" required margin="none">
                        <FileInput
                            name="files"
                            renderFile={(file) => (
                                <Box pad="small" direction="row" justify="between" align="center">
                                    <Box width="xsmall" height="xsmall">
                                        <Image src={ URL.createObjectURL(file) } fit="cover" />
                                    </Box>
                                    <Text>{file.name}</Text>
                                </Box>
                            )}
                            multiple={{"aggregateThreshold": 5}}
                            messages={{
                                browse: "выбрать",
                                dropPrompt: "перетащите файл сюда",
                                dropPromptMultiple: "перетащите файлы сюда",
                                files: "файлов",
                                remove: "удалить",
                                removeAll: "удалить все"
                            }}
                            onChange={handleFileInputChange}
                        />
                    </FormField>
                    <Box align="center">
                        <Button
                            type="submit"
                            primary
                            label={loading ? 'Загрузка...' : 'Загрузить' }
                            disabled={loading}
                        />
                    </Box>
                </Box>
            </Form>
        </Layer>
    )
}