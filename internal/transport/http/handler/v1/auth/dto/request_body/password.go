package request_body

import "github.com/mephistolie/chefbook-backend-api-gateway/internal/transport/http/contract"

type RequestPasswordReset = contract.RequestPasswordResetRequest
type ResetPassword = contract.ConfirmPasswordResetRequest
type ChangePassword = contract.SetPasswordRequest
