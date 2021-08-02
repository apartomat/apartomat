import React from "react";

function CheckEmail({ email }: { email: string }) {
    return (
        <div>
            <h1>Verifcation code sent</h1>
            <p>Please check email {email} and go to by link.</p>
        </div>
    );
}

export default CheckEmail;