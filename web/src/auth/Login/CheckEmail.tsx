import React from "react";

import Box from "@material-ui/core/Box";

function CheckEmail({ email }: { email: string }) {
    return (
        <Box width={1/4}>
            <h1>Verifcation code sent</h1>
            <p>Please check email {email} and go to by link.</p>
        </Box>
    );
}

export default CheckEmail;