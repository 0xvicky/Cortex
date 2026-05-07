import { BrowserRouter, Link, Navigate, Route, Routes, useLocation } from 'react-router-dom'
import { GoogleOAuthProvider } from '@react-oauth/google'
import HomePage from './pages/HomePage'
import JobsPage from './pages/JobsPage'
import JobDetailPage from './pages/JobDetailPage'
import AuthPage from './pages/AuthPage'
import { AuthProvider, useAuth } from './context/AuthContext'
import { ProtectedRoute } from './components/ProtectedRoute'

function ActiveNav({ to, children }: { to: string; children: React.ReactNode }) {
  const location = useLocation()
  const isActive = location.pathname === to
  return (
    <Link
      to={to}
      className={`px-3 py-2 text-xl font-medium transition ${isActive ? 'text-purple-400' : 'text-gray-400 hover:text-white'}`}
    >
      {children}
    </Link>
  )
}

function AppShell() {
  const { isAuthenticated, user, logout } = useAuth()
  const location = useLocation()

  // Don't show header on auth page
  if (location.pathname === '/auth') {
    return (
      <main className="flex-1">
        <Routes>
          <Route path="/auth" element={<AuthPage />} />
          <Route path="*" element={<Navigate to="/auth" replace />} />
        </Routes>
      </main>
    )
  }

  return (
    <div className="relative min-h-screen text-white">
      <div className="mx-auto flex min-h-screen max-w-6xl flex-col px-4 py-8 sm:px-6">
        <header className="mb-10 flex items-center gap-6 border-b border-white/10 pb-4">
          <Link to="/" className="text-3xl font-semibold text-white hover:text-purple-400 transition font-mono">
            Cortex
          </Link>
          <nav className="ml-auto flex items-center gap-4">
            <ActiveNav to="/">Home</ActiveNav>
            <ActiveNav to="/jobs">Jobs</ActiveNav>
          </nav>
          {isAuthenticated && (
            <div className="flex items-center gap-4 border-l border-white/10 pl-4">
              <div className="flex items-center gap-2">
                {user?.picture && (
                  <img src={user.picture} alt={user.name} className="w-8 h-8 rounded-full" />
                )}
                <span className="text-sm text-gray-300">{user?.name}</span>
              </div>
              <button
                onClick={logout}
                className="px-3 py-2 text-sm font-medium text-gray-400 hover:text-white transition rounded hover:bg-white/10"
              >
                Logout
              </button>
            </div>
          )}
        </header>

        <main className="flex-1">
          <Routes>
            <Route path="/auth" element={<AuthPage />} />
            <Route
              path="/"
              element={
                <ProtectedRoute>
                  <HomePage />
                </ProtectedRoute>
              }
            />
            <Route
              path="/jobs"
              element={
                <ProtectedRoute>
                  <JobsPage />
                </ProtectedRoute>
              }
            />
            <Route
              path="/jobs/:userId/:job_id"
              element={
                <ProtectedRoute>
                  <JobDetailPage />
                </ProtectedRoute>
              }
            />
            <Route path="*" element={<Navigate to="/" replace />} />
          </Routes>
        </main>
      </div>
    </div>
  )
}

export default function App() {
  const googleClientId = import.meta.env.VITE_GOOGLE_CLIENT_ID

  if (!googleClientId) {
    return (
      <div className="flex items-center justify-center min-h-screen bg-slate-900 text-white">
        <div className="text-center">
          <h1 className="text-2xl font-bold mb-2">Configuration Error</h1>
          <p className="text-gray-400">
            Please set VITE_GOOGLE_CLIENT_ID in your .env file
          </p>
        </div>
      </div>
    )
  }

  return (
    <GoogleOAuthProvider clientId={googleClientId}>
      <AuthProvider>
        <BrowserRouter>
          <AppShell />
        </BrowserRouter>
      </AuthProvider>
    </GoogleOAuthProvider>
  )
}
