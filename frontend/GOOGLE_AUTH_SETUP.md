# Google Authentication Setup Guide

## Overview

This guide will help you set up Google OAuth 2.0 authentication for the Cortex frontend.

## Step 1: Create a Google Cloud Project

1. Go to [Google Cloud Console](https://console.cloud.google.com/)
2. Click on the project dropdown and select "NEW PROJECT"
3. Enter a project name (e.g., "Cortex") and click "CREATE"
4. Wait for the project to be created

## Step 2: Enable Google+ API

1. In the Cloud Console, go to "APIs & Services" → "Library"
2. Search for "Google+ API"
3. Click on it and then click "ENABLE"

## Step 3: Create OAuth 2.0 Credentials

1. Go to "APIs & Services" → "Credentials"
2. Click "CREATE CREDENTIALS" → "OAuth Client ID"
3. If prompted, configure the OAuth consent screen first:
   - Choose "External" for User Type
   - Fill in required fields (App name, User support email, etc.)
   - Add your email as a test user
4. For the OAuth Client ID creation:
   - Application type: **Web application**
   - Name: "Cortex Frontend"
   - Authorized JavaScript origins:
     - `http://localhost:5173` (for development)
     - `http://localhost:5173/` (with trailing slash)
     - Your production domain (e.g., `https://cortex.example.com`)
   - Authorized redirect URIs:
     - `http://localhost:5173/auth`
     - Your production redirect URI

5. Click "CREATE"
6. Copy the **Client ID** (you'll need this next)

## Step 4: Configure Environment Variables

1. In the `frontend` directory, create or update `.env.local`:

   ```bash
   VITE_GOOGLE_CLIENT_ID=your_client_id_here
   VITE_GITHUB_TOKEN=your_github_token_here
   ```

2. Replace `your_client_id_here` with the Client ID you copied from Google Cloud Console

## Step 5: Run the Development Server

```bash
cd frontend
npm run dev
```

The app will start at `http://localhost:5173` and automatically redirect to the login page.

## Features

### Authentication Flow

- Users are redirected to `/auth` if not authenticated
- Google login button allows users to sign in with their Google account
- User data (name, email, picture) is stored in localStorage
- User profile appears in the header with logout option

### Protected Routes

- All routes except `/auth` are protected
- Unauthenticated users are redirected to the login page
- User session persists across page reloads

### User Profile Display

- User's name and profile picture appear in the header
- Logout button removes authentication and redirects to login

## Troubleshooting

### "Configuration Error: Please set VITE_GOOGLE_CLIENT_ID"

- Ensure your `.env.local` file has the `VITE_GOOGLE_CLIENT_ID` variable set
- Restart the development server after updating `.env.local`

### CORS or Redirect URI Errors

- Verify you've added `http://localhost:5173` to authorized JavaScript origins
- Verify you've added `http://localhost:5173/auth` to authorized redirect URIs
- The URLs must match exactly (including protocol and port)

### Google Login Button Not Showing

- Check browser console for errors
- Ensure the Client ID is valid and correct
- Clear browser cache and reload

## File Structure

```
frontend/src/
├── App.tsx                          # Main app with OAuth provider
├── context/
│   └── AuthContext.tsx             # Authentication state management
├── components/
│   └── ProtectedRoute.tsx           # Route protection wrapper
└── pages/
    └── AuthPage.tsx                # Google login page
```

## Environment Variables Reference

| Variable                | Description                 | Example                                |
| ----------------------- | --------------------------- | -------------------------------------- |
| `VITE_GOOGLE_CLIENT_ID` | Google OAuth Client ID      | `123456789.apps.googleusercontent.com` |
| `VITE_GITHUB_TOKEN`     | GitHub API token (existing) | Your GitHub personal access token      |

## Security Notes

- Never commit `.env.local` to version control
- Keep your Google Client ID secure
- Use environment-specific credentials for development and production
- Consider implementing token refresh logic for production

## Next Steps

1. Set up backend authentication endpoints to validate Google tokens
2. Implement session management on the backend
3. Add role-based access control if needed
4. Set up production OAuth credentials with your domain
