import React, { useState } from "react";

import { LoginByEmailFn } from "./useLoginByEmail";

import Box from "@material-ui/core/Box";
import FormControl from "@material-ui/core/FormControl";
import FormHelperText from "@material-ui/core/FormHelperText";
import Input from "@material-ui/core/Input";
import Button from "@material-ui/core/Button";

function Form( { loading, login, error }:
    { loading: boolean, login: LoginByEmailFn, error?: { message: string } }
) {
    const [email, setEmail] = useState("");
    const [emailInputError, setEmailInputError] = useState("");

    function handleInput(event: React.ChangeEvent<HTMLInputElement>) {
        setEmail(event.target.value);
    }

    function handleSubmit(event: React.FormEvent) {
        if (email === "") {
            setEmailInputError("Email is required!");

        } else {
            setEmailInputError("");
            login(email);
        }

        event.preventDefault();
    }

    return (
        <Box width={1/4}>
            <h1>Sign in</h1>
            {error ? <div>{error.message}</div> : null}
            <form onSubmit={handleSubmit}>
                <FormControl fullWidth margin="normal">
                    <Input value={email} onChange={handleInput} placeholder="Email *" error={emailInputError !== ""}/>
                    {emailInputError ? <FormHelperText>{emailInputError}</FormHelperText> : ""}
                </FormControl>
                <FormControl margin="normal">
                    <Button type="submit" variant="contained" disableRipple={true} color="secondary" disabled={loading === true}>Sign in</Button>
                </FormControl>
            </form>
        </Box>
    );
}

export default Form;