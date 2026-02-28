import { BrowserRouter, Routes, Route, Navigate } from "react-router-dom";
import Login from "./pages/Login";
import Signup from "./pages/Signup";
import Chat from "./pages/Chat";
import Rooms from "./pages/Rooms";
import Profile from "./pages/Profile";
import AppShell from "./components/AppShell";
import { ToastProvider } from "./components/Toast";

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
