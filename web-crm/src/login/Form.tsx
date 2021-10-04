import React, { useState } from "react";

import { LoginByEmailFn } from "./useLoginByEmail";

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
        <div>
            <h1>Sign in</h1>
            {error ? <div>{error.message}</div> : null}
            <form onSubmit={handleSubmit}>
                <p>
                    <input value={email} onChange={handleInput} placeholder="Email *"/>
                </p>
                {emailInputError ? <p>{emailInputError}</p> : null}
                <p>
                    <button type="submit" disabled={loading === true}>Sign in</button>
                </p>
            </form>
        </div>
    );
}

export default Form;