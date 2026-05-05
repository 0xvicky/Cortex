import { BrowserRouter, Link, Navigate, Route, Routes, useLocation } from 'react-router-dom'
import HomePage from './pages/HomePage'
import JobsPage from './pages/JobsPage'
import JobDetailPage from './pages/JobDetailPage'

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
        </header>

        <main className="flex-1">
          <Routes>
            <Route path="/" element={<HomePage />} />
            <Route path="/jobs" element={<JobsPage />} />
            <Route path="/jobs/:userId/:job_id" element={<JobDetailPage />} />
            <Route path="*" element={<Navigate to="/" replace />} />
          </Routes>
        </main>
      </div>
    </div>
  )
}

export default function App() {
  return (
    <BrowserRouter>
      <AppShell />
    </BrowserRouter>
  )
}
