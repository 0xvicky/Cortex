import { useState } from 'react'
import { useNavigate } from 'react-router-dom'
import { createJob } from '../api'
import { Octokit } from "@octokit/rest";
// const repoPattern = /^https:\/\/github\.com\/([^/]+)\/([^/]+)(?:\/)?$/
import { getOwner } from '../utils/utils';
const octokit = new Octokit({ auth: import.meta.env.VITE_GITHUB_TOKEN });


// console.log(octokit)

export default function HomePage() {
  const [repoUrl, setRepoUrl] = useState('')
  const [touched, setTouched] = useState(false)
  const [error, setError] = useState<string | null>(null)
  const [loading, setLoading] = useState(false)
  const [checking, setChecking] = useState(false)
  const [validRepo, setValidRepo] = useState(false)
  const navigate = useNavigate()
  const userId = "whyvickyyy";

  const trimmedRepoUrl = repoUrl.trim()
  const showError = touched && trimmedRepoUrl.length > 0 && !validRepo && !checking
  const helperText = showError
    ? 'Enter a valid GitHub repo URL'
    : 'Example: https://github.com/vercel/next.js'


  const validateRepo = async (url: string) => {
    setChecking(true)
    setValidRepo(false)

    if (!url.trim()) {
      setChecking(false)
      return
    }

    try {
        const{parts, isGithub} = getOwner(url);
      
      if (!isGithub || parts.length < 2) {
        setValidRepo(false)
        return
      }

      const [owner, repo] = parts
      await octokit.repos.get({ owner, repo })
      setValidRepo(true)
    } catch (err) {
      setValidRepo(false)   
    } finally {
      setChecking(false)
    }
  }
  const handleSubmit = async (event: React.FormEvent) => {
    event.preventDefault()
    setTouched(true)
    setError(null)

    if (!validRepo) {
      return
    }

    try {
      setLoading(true)
      const result = await createJob(trimmedRepoUrl,userId)
      navigate(`/jobs/${userId}/${result.job_id}`)
    } catch (err) {
      setError(err instanceof Error ? err.message : 'Unable to analyze repo')
    } finally {
      setLoading(false)
    }
  }

  const inputStateClass = showError
    ? 'border-rose-500/60 text-white placeholder:text-rose-300 focus:border-purple-400 focus:ring-purple-500/40'
    : 'border-white/10 text-white placeholder:text-slate-500 focus:border-purple-400 focus:ring-purple-500/40'

  return (
    <section className="mx-auto flex w-full max-w-3xl flex-col gap-10 rounded-xl border border-white/10 bg-white/5 p-10 shadow-[0_40px_120px_-80px_rgba(124,58,237,0.6)] backdrop-blur-sm">
      <div className="space-y-5">
        <h1 className="text-5xl font-semibold leading-tight text-white">Analyze a repository in seconds.</h1>
        <p className="max-w-2xl text-gray-400 leading-8">
          Paste a GitHub repo URL and get a concise summary, plus an interactive chat for follow-up questions.
        </p>
      </div>

      <form className="grid gap-5" onSubmit={handleSubmit} noValidate>
        <label className="space-y-6 text-sm font-medium text-white" htmlFor="repo-url">
          <p className="font-bold">Repository URL</p>
          <div className="relative">
            <input
              id="repo-url"
              type="url"
              value={repoUrl}
              onChange={(e) => {
                setRepoUrl(e.target.value)
                validateRepo(e.target.value)
              }}
              onBlur={() => setTouched(true)}
              placeholder="https://github.com/owner/repo"
              className={`w-full rounded-lg border bg-slate-950/95 px-6 py-4 pr-16 text-base outline-none transition duration-200 ${inputStateClass}`}
              autoComplete="off"
              spellCheck="false"
            />
            {trimmedRepoUrl.length > 0 ? (
              <span className="pointer-events-none absolute inset-y-0 right-4 flex items-center justify-center">
                {checking ? (
                  <span className="inline-flex h-5 w-5 animate-spin rounded-full border-2 border-white/10 border-t-purple-400" />
                ) : validRepo ? (
                  <span className="inline-flex h-9 w-9 items-center justify-center rounded-full border border-emerald-400 bg-emerald-500/10 text-emerald-400">
                    ✓
                  </span>
                ) : (
                  <span className="inline-flex h-9 w-9 items-center justify-center rounded-full border border-rose-400 bg-rose-500/10 text-rose-400">
                    ✕
                  </span>
                )}
              </span>
            ) : null}
          </div>
        </label>
        <div className="flex flex-col gap-4 sm:flex-row sm:items-center sm:justify-between">
          <p className={`text-sm ${showError ? 'text-rose-300' : 'text-gray-400'}`}>{helperText}</p>
          <button
            type="submit"
            disabled={!validRepo || loading}
            className="group relative inline-flex cursor-pointer items-center justify-center rounded-lg bg-gradient-to-r from-purple-600 to-violet-600 px-8 py-3 text-sm font-semibold text-white shadow-[0_18px_50px_-30px_rgba(124,58,237,0.9)] transition duration-200 hover:shadow-[0_22px_80px_-40px_rgba(124,58,237,0.8)] hover:ring-2 hover:ring-purple-400/50 disabled:cursor-not-allowed disabled:bg-slate-700 disabled:shadow-none disabled:ring-0"
          >
            {loading ? 'Analyzing...' : 'Send'}
          </button>
        </div>
        {error ? <p className="text-sm text-rose-300">{error}</p> : null}
      </form>
    </section>
  )
}
