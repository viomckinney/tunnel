import { Route } from "react-router";
import { BrowserRouter, Routes } from "react-router-dom";
import DashboardPage from "./pages/dashboard";
import IndexPage from "./pages/index_page";
import LoginPage from "./pages/login";
import RegisterPage from "./pages/register";

function App() {
    return (
        <>
            <BrowserRouter>
                <Routes>
                    <Route path="/">
                        <Route index element={<IndexPage />} />
                        <Route path="/login" element={<LoginPage />} />
                        <Route path="/register" element={<RegisterPage />} />
                        <Route path="/dashboard" element={<DashboardPage />} />
                    </Route>
                </Routes>
            </BrowserRouter>
        </>
    );
}

export default App;
