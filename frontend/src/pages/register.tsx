import { useState } from "react";
import { Navigate } from "react-router";

export default function RegisterPage() {
    const [username, setUsername] = useState("");
    const [password, setPassword] = useState("");
    const [passwordConfirm, setPasswordConfirm] = useState("");
    const [inviteCode, setInviteCode] = useState("");
    const [redirect, setRedirect] = useState("");
    const [error, setError] = useState("");

    function go() {
        if (password !== passwordConfirm) {
            setError("Passwords do not match");
            return;
        }

        fetch("/api/register", {
            method: "POST",
            body: JSON.stringify({ username, password, inviteCode }),
        }).then(async (res) => {
            const text = await res.text();

            if (text === "OK") {
                setRedirect("/dashboard");
            } else {
                setError(text);
            }
        });
    }

    return (
        <>
            <h1>login</h1>
            <div>
                username:
                <input
                    type="text"
                    placeholder="jdoe"
                    onChange={(e) => setUsername(e.target.value)}
                />
            </div>
            <div>
                password:
                <input
                    type="password"
                    placeholder="hunter2"
                    onChange={(e) => setPassword(e.target.value)}
                />
            </div>
            <div>
                confirm password:
                <input
                    type="password"
                    placeholder="hunter2"
                    onChange={(e) => setPasswordConfirm(e.target.value)}
                />
            </div>
            <div>
                invite code:
                <input
                    type="text"
                    placeholder="super-cool-invite-code"
                    onChange={(e) => setInviteCode(e.target.value)}
                />
            </div>
            <div>
                <button
                    onClick={(e) => {
                        e.preventDefault();
                        go();
                    }}
                >
                    go
                </button>
            </div>
            <div>
                <p style={{ color: "red" }}>{error}</p>
            </div>
            {redirect && <Navigate to={redirect} />}
        </>
    );
}
