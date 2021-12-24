import React from "react"

import { Main, Heading, Paragraph } from "grommet"

export default function Error({ message }: { message: string }) {
    return (
        <Main pad="large">
            <Heading>Ошибка</Heading>
            <Paragraph>{message}</Paragraph>
        </Main>
    );
}