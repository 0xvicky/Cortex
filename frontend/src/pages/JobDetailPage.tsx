import { useEffect, useMemo, useRef, useState } from 'react'
import { useNavigate, useParams, useLocation } from 'react-router-dom'
import { fetchSummary, queryJob } from '../api'
import type { RepoSummary, RepoOwner } from '../types'
import { getOwner } from '../utils/utils'
interface ChatMessage {
  id: string
  role: 'user' | 'ai'
  text: string
}

const repoPattern = /^https:\/\/github\.com\/([^/]+)\/([^/]+)(?:\/)?$/

function formatRepoName(repoUrl: string) {
  const match = repoPattern.exec(repoUrl)
  return match ? `${match[1]}/${match[2]}` : repoUrl
}

function typingDots() {
  return (
    <span className="inline-flex items-center gap-2 text-gray-400">
      <span className="h-2.5 w-2.5 animate-pulse rounded-full bg-purple-400" />
      <span className="h-2.5 w-2.5 animate-pulse rounded-full bg-purple-400 delay-75" />
      <span className="h-2.5 w-2.5 animate-pulse rounded-full bg-purple-400 delay-150" />
    </span>
  )
}

export default function JobDetailPage() {
//   const { job_id  } = useParams()
  const navigate = useNavigate()
  const {userId,job_id} = useParams()
  const [summary, setSummary] = useState<RepoSummary | null>(null)
  const [status, setStatus] = useState<string>('pending')
  const [summaryError, setSummaryError] = useState<string | null>(null)
  const [summaryLoading, setSummaryLoading] = useState(true)
  const [messages, setMessages] = useState<ChatMessage[]>([])
  const [question, setQuestion] = useState('')
  const [sending, setSending] = useState(false)
  const [chatError, setChatError] = useState<string | null>(null)
  const messagesRef = useRef<HTMLDivElement | null>(null)
  const [repoOwner, setRepoOwner] = useState<RepoOwner | null>(null)
  const location = useLocation()

  const repoName = `${repoOwner?.owner}/${repoOwner?.repoName}`
//   console.log(repoName)
// const repoName = 

// console.log(userId)
  useEffect(() => {
    if (!job_id) return

    let active = true
    let intervalId = 0

    const loadSummary = async () => {
      try {
        const data = await fetchSummary(job_id,userId)
        // console.log("in detail page")
        // console.log(data)
        // console.log(data)
        const {parts} = getOwner(data?.job.repo_url);
        const [owner, repo] = parts
        if (!active) return
        
        if (data.job?.repo_summary) {
            setRepoOwner({
              owner, repoName:repo
            })
          setSummary(data.job.repo_summary)
          setStatus(data?.job?.status?? 'PENDING')
          setSummaryLoading(false)
          if (data?.job?.status === 'COMPLETED'){
              setStatus(data?.job?.status)
            clearInterval(intervalId)
            return
          }
        
        } else {
          setSummary(null)
          setStatus(data.job?.status || 'pending')
          setSummaryLoading(false)
        }
      } catch (err) {
        if (!active) return
        setSummaryError(err instanceof Error ? err.message : 'Unable to load summary')
        setStatus("failed")
        setSummaryLoading(false)
        clearInterval(intervalId)
      }
    }

    loadSummary()
    intervalId = window.setInterval(loadSummary, 3000)

    return () => {
        active = false
        clearInterval(intervalId)
        }
    }, [job_id])

    useEffect(() => {
        if (!messagesRef.current) return
        messagesRef.current.scrollTop = messagesRef.current.scrollHeight
    }, [messages, sending])

    const handleSend = async () => {
        if (!question.trim() || sending || !job_id || status !== "COMPLETED") return

        const text = question.trim()
        setMessages((current) => [...current, { id: `user-${Date.now()}`, role: 'user', text }])
        setQuestion('')
        setChatError(null)
        setSending(true)

        try {
        const data = await queryJob(job_id, text)
        console.log(data)
        setMessages((current) => [...current, { id: `ai-${Date.now()}`, role: 'ai', text: data.res }])
        } catch (err) {
        setChatError(err instanceof Error ? err.message : 'Unable to send question')
        } finally {
        setSending(false)
        }
    }

    const handleKeyDown = (event: React.KeyboardEvent<HTMLTextAreaElement>) => {
    if (event.key === 'Enter' && !event.shiftKey) {
      event.preventDefault()
      handleSend()
    }
  }

  if (!job_id) {
    return (
      <div className="rounded-[2rem] border border-white/10 bg-white/5 p-4 sm:p-10 text-white backdrop-blur-sm">
        <p className="text-xl sm:text-2xl font-semibold">Missing job identifier.</p>
        <button
          type="button"
          onClick={() => navigate('/jobs')}
          className="mt-6 rounded-[1.5rem] bg-gradient-to-r from-purple-600 to-violet-600 px-6 py-3 text-sm font-semibold text-white shadow-[0_18px_50px_-30px_rgba(124,58,237,0.9)]"
        >
          Back to jobs
        </button>
      </div>
    )
  }

  return (
    <section className="flex min-h-[calc(100vh-4rem)] flex-col gap-10 ">
      <div className="flex flex-col gap-6 sm:flex-row sm:items-center sm:justify-between">
        <div className="space-y-1 max-w-2xl min-w-0">
          <h1 className="text-2xl sm:text-4xl font-semibold leading-tight text-white break-words">Repository analysis</h1>
          <p className="text-base leading-7 text-gray-400 break-words">Chat with the indexed repository once the summary is ready.</p>
        </div>
        <div className="w-full sm:w-auto rounded-[1.5rem] border border-white/10 bg-slate-950/80 px-5 py-4 text-sm text-gray-300 shadow-inner shadow-black/30 overflow-hidden">
          <p className="font-medium text-white break-words">Job ID {job_id.slice(0, 8)}</p>
          <p className="mt-1 text-gray-400 break-words">Status {status}</p>
        </div>
      </div>

      <div className="grid grid-cols-1 gap-8 xl:grid-cols-[1.2fr_0.8fr]">
        <div className="rounded-[2rem] border border-white/10 bg-white/5 p-4 sm:p-8 backdrop-blur-sm shadow-[0_24px_80px_-40px_rgba(124,58,237,0.45)]">
        {summary && <h2 className="mb-6 text-2xl sm:text-3xl font-semibold text-white break-words">{repoName}</h2>}  
          <div className="rounded-[1.75rem] border border-white/10 bg-slate-950/90 p-4 sm:p-8 overflow-hidden">
            {summaryLoading ? (
              <div className="space-y-5">
                <div className="h-5 w-3/4 animate-pulse rounded-full bg-slate-800" />
                <div className="h-5 w-full animate-pulse rounded-full bg-slate-800" />
                <div className="h-5 w-5/6 animate-pulse rounded-full bg-slate-800" />
                <div className="h-5 w-11/12 animate-pulse rounded-full bg-slate-800" />
              </div>
            ) : summaryError ? (
              <div className="space-y-4 text-rose-200">
                <p className="text-xl font-semibold text-white">Unable to load summary</p>
                <p className="text-sm leading-7 text-gray-300">{summaryError}</p>
              </div>
            ) : summary ? (
              <div className="space-y-8 overflow-hidden">
                {summary.project_purpose ? (
                  <div className="space-y-3">
                    <p className="text-sm uppercase tracking-[0.3em] text-gray-400">Project Purpose</p>
                    <p className="text-base leading-8 text-gray-100 break-words">{summary.project_purpose}</p>
                  </div>
                ) : null}
                {summary.tech_stack ? (
                  <div className="space-y-3">
                    <p className="text-sm uppercase tracking-[0.3em] text-gray-400">Tech Stack</p>
                    <div className="flex flex-wrap gap-2">
                      {summary.tech_stack.map((tech) => (
                        <span key={tech} className="rounded-full border border-white/10 bg-slate-900 px-3 py-1 text-sm text-gray-200">
                          {tech}
                        </span>
                      ))}
                    </div>
                  </div>
                ) : null}
                {summary.key_components ? (
                  <div className="space-y-3">
                    <p className="text-sm uppercase tracking-[0.3em] text-gray-400">Key Components</p>
                    <ul className="list-disc space-y-2 pl-5 text-gray-200 flex flex-col overflow-hidden">
                      {summary.key_components.map((item, index) => {
                        if (typeof item === 'string') {
                          return <li key={index} className="break-words">{item}</li>
                        }
                        return (
                          <li key={index} className="break-words">
                            <span className="font-semibold text-white">{item.name ?? 'Component'}:</span>{' '}
                            <span className="break-words">{item.description ?? JSON.stringify(item)}</span>
                          </li>
                        )
                      })}
                    </ul>
                  </div>
                ) : null}
                {summary.architecture ? (
                  <div className="space-y-3">
                    <p className="text-sm uppercase tracking-[0.3em] text-gray-400">Architecture</p>
                    <p className="text-base leading-8 text-gray-100 break-words">{summary.architecture}</p>
                  </div>
                ) : null}
                {summary.use_cases ? (
                  <div className="space-y-3">
                    <p className="text-sm uppercase tracking-[0.3em] text-gray-400">Use Cases</p>
                    <ul className="list-disc space-y-2 pl-5 text-gray-200 overflow-hidden">
                      {summary.use_cases.map((useCase) => (
                        <li key={useCase} className="break-words">{useCase}</li>
                      ))}
                    </ul>
                  </div>
                ) : null}
                {summary.how_to_run ? (
                  <div className="space-y-3">
                    <p className="text-sm uppercase tracking-[0.3em] text-gray-400">How to Run</p>
                    {typeof summary.how_to_run === 'string' ? (
                      <p className="text-base leading-8 text-gray-100 break-words">{summary.how_to_run}</p>
                    ) : (
                      <ul className="list-disc space-y-2 pl-5 text-gray-200 overflow-hidden">
                        {summary.how_to_run.map((step, index) => (
                          <li key={`${step}-${index}`} className="break-words">{step}</li>
                        ))}
                      </ul>
                    )}
                  </div>
                ) : null}
              </div>
            ) : (
              <div className="space-y-4 text-gray-400">
                <p className="text-xl font-semibold text-white">Waiting for analysis...</p>
                <p className="text-sm leading-7">The repo ingestion is still in progress. This panel will refresh automatically.</p>
              </div>
            )}
          </div>
        </div>

          <div className="rounded-[2rem] border border-white/10 bg-white/5 p-4 sm:p-8 backdrop-blur-sm shadow-[0_24px_80px_-40px_rgba(124,58,237,0.45)] overflow-hidden">
          <div className="mb-6 flex items-center justify-between gap-4 overflow-hidden">
            <h2 className="text-2xl sm:text-3xl font-semibold text-white truncate">Ask about the repo</h2>
            {sending ? typingDots() : null}
          </div>
          <div ref={messagesRef} className="mb-6 max-h-[280px] sm:max-h-[420px] space-y-4 overflow-y-auto overflow-x-hidden rounded-[1.75rem] border border-white/10 bg-slate-950/90 p-5">
            {messages.length === 0 ? (
              <div className="rounded-[1.5rem] bg-slate-900/80 p-5 text-gray-400">Start by asking a question about the repository.</div>
            ) : (
              messages.map((message) => (
                <div
                  key={message.id}
                  className={`rounded-[1.75rem] border p-4 text-sm leading-7 overflow-hidden ${message.role === 'user'
                    ? 'ml-auto max-w-[85%] border-white/10 bg-violet-500/10 text-white shadow-[0_20px_60px_-40px_rgba(124,58,237,0.6)]'
                    : 'mr-auto max-w-[85%] border-white/10 bg-slate-900 text-gray-100'}`}
                >
                  {message.role === 'user' ? <p className="text-xs uppercase tracking-[0.25em] text-purple-300">You</p> : <p className="text-xs uppercase tracking-[0.25em] text-gray-400">AI</p>}
                  <p className="mt-3 whitespace-pre-wrap break-words">{message.text}</p>
                </div>
              ))
            )}
          </div>

          <div className="space-y-4 rounded-[1.75rem] border border-white/10 bg-slate-950/90 p-4 sm:p-5">
            <label className="block text-sm font-medium text-gray-200" htmlFor="question">
              Ask a question
            </label>
            <textarea
              id="question"
              value={question}
              onChange={(event) => setQuestion(event.target.value)}
              onKeyDown={handleKeyDown}
              rows={3}
              className="w-full rounded-[1.5rem] border border-white/10 bg-slate-900/95 px-5 py-4 text-base text-white outline-none transition duration-200 focus:border-purple-400 focus:ring-2 focus:ring-purple-500/40"
              placeholder="What should I know about this repository?"
            />
            <div className="flex flex-col gap-3 sm:flex-row sm:items-center sm:justify-between">
              <p className="text-xs text-gray-400">Press Enter to send. Shift+Enter for new line.</p>
              <button
                type="button"
                onClick={handleSend}
                disabled={!question.trim() || sending || status=="PENDING" || status=="FAILED" }
                className="inline-flex items-center justify-center rounded-[1.5rem] bg-gradient-to-r from-purple-600 to-violet-600 px-6 py-3 text-sm font-semibold text-white shadow-[0_18px_50px_-30px_rgba(124,58,237,0.9)] transition duration-200 hover:shadow-[0_22px_80px_-40px_rgba(124,58,237,0.8)] cursor-pointer disabled:cursor-not-allowed disabled:bg-slate-700 disabled:shadow-none"
              >
              {status == "PENDING" ? 
              <>
                <svg className="mr-2 h-8 w-8 animate-spin text-white" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
                  <circle className="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" strokeWidth="4"></circle>
                  <path className="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4z"></path>
                </svg>
                Embedding
              </>
              : sending ? 'Sending...' : 'Ask question'}
              </button>
            </div>
            {chatError ? <p className="text-sm text-rose-300">{chatError}</p> : null}
          </div>
        </div>
      </div>
    </section>
  )
}
