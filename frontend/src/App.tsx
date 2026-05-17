import { BrowserRouter, Navigate, Route, Routes } from "react-router-dom";
import AppShell from "./components/AppShell";
import { ToastProvider } from "./components/Toast";
import Chat from "./pages/Chat";
import Login from "./pages/Login";
import Profile from "./pages/Profile";
import Rooms from "./pages/Rooms";
import Signup from "./pages/Signup";

function App() {
  return (
    <ToastProvider>
      <BrowserRouter>
        <div className="min-h-screen bg-[#f6f1e7] text-[#2b2620]">
          <Routes>
            <Route element={<AppShell />}>
              <Route path="/chat" element={<Chat />} />
              <Route path="/rooms" element={<Rooms />} />
              <Route path="/profile" element={<Profile />} />
            </Route>
            <Route path="/login" element={<Login />} />
            <Route path="/signup" element={<Signup />} />
            <Route path="*" element={<Navigate to="/login" replace />} />
          </Routes>
        </div>
      </BrowserRouter>
    </ToastProvider>
  );
}

export default App;
