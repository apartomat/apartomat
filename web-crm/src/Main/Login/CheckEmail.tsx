import React from "react"

import { Main, Heading, Box } from "grommet"

function CheckEmail({ email }: { email: string }) {
    return (
        <Main pad="large">
            <Box width="large">
                <Heading level={2}>Код отправлен</Heading>
                <Box>Для входа проверьте почту {email} и пройди по ссылке.</Box>
            </Box>
        </Main>
    )
}

export default CheckEmail