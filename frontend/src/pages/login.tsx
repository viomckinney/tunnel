import { useState } from "react";
import { Navigate } from "react-router";

export default function LoginPage() {
    const [username, setUsername] = useState("");
    const [password, setPassword] = useState("");
    const [redirect, setRedirect] = useState("");
    const [error, setError] = useState("");

    function go() {
        fetch("/api/login", {
            method: "POST",
            body: JSON.stringify({ username, password }),
        }).then(async (res) => {
            const text = await res.text();

            if (!res.ok) {
                setError(text);
            } else {
                localStorage.token = text;
                setRedirect("/dashboard");
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
