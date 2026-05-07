import { GoogleLogin, type CredentialResponse } from '@react-oauth/google'
import { useNavigate } from 'react-router-dom'
import { useAuth } from '../context/AuthContext'
import { jwtDecode } from 'jwt-decode'
import {handleGoogleAuth} from '../api'
interface GoogleJWT {
  email: string
  name: string
  picture?: string
  sub: string
}

export default function AuthPage() {
  const navigate = useNavigate()
  const { login } = useAuth()

  const handleSuccess = async(credentialResponse: CredentialResponse) => {
    if (!credentialResponse.credential) return

    try {
      const {token} =await handleGoogleAuth(credentialResponse.credential)
      console.log(token)
      const decoded = jwtDecode<GoogleJWT>(credentialResponse.credential)
      login({
        id: decoded.sub,
        email: decoded.email,
        name: decoded.name,
        picture: decoded.picture,
        jwt:token
      })

      navigate('/')
    } catch (error) {
      console.error('Failed to decode token:', error)
    }
  }

  return (
    <div className="min-h-screen bg-gradient-to-br from-slate-900 via-purple-900 to-slate-900 flex items-center justify-center px-4">
      <div className="max-w-md w-full">
        <div className="bg-slate-800/50 backdrop-blur-sm border border-white/10 rounded-lg p-8 text-center">
          <h1 className="text-4xl font-bold font-mono text-white mb-2">Cortex</h1>
          <p className="text-gray-400 mb-8">Sign in to continue</p>

          <div className="flex justify-center">
            <GoogleLogin
              onSuccess={handleSuccess}
              onError={() => console.log('Login Failed')}
            />
          </div>

          <p className="text-xs text-gray-500 mt-6">
            By signing in, you agree to our Terms of Service and Privacy Policy
          </p>
        </div>
      </div>
    </div>
  )
}
