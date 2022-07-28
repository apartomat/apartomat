import React, { useState, useRef, Dispatch, SetStateAction} from "react"

import { Box, Button, Drop, Text } from "grommet"
import { Next } from "grommet-icons"

export default function AddSomething ({ showUploadFiles }: { showUploadFiles: Dispatch<SetStateAction<boolean>> }) {
    const [show, setShow] = useState(false)

    const targetRef = useRef<HTMLDivElement>(null)

    return (
        <Box justify="center">
            <Box ref={targetRef}>
                <Button label="Добавить" icon={<Next/>} reverse onClick={() => setShow(true)} />
            </Box>

            {show && targetRef.current && (
                <Drop
                    elevation="small"
                    round="small"
                    align={{top: "bottom", right: "right"}}
                    margin={{top: "xsmall"}}
                    target={targetRef.current}
                    onClickOutside={() => setShow(false)}
                    onEsc={() => setShow(false)}
                >
                    <Button plain children={({ hover }: {hover: boolean}) => {
                        return <Box pad="small" background={hover ? 'light-1': ''}><Text>Визуализации</Text></Box>
                    }} onClick={() => {
                        showUploadFiles(true)
                        setShow(false)
                    }}/>
                    <Button plain>
                        <Box pad="small"><Text>План</Text></Box>
                    </Button>
                    <Button plain>
                        <Box pad="small"><Text>Исходники</Text></Box>
                    </Button>
                    <Button plain>
                        <Box pad="small"><Text>Альбом</Text></Box>
                    </Button>
                    <Button plain>
                        <Box pad="small"><Text>Спецификация</Text></Box>
                    </Button>
                </Drop>
            )}
        </Box>
    )
}