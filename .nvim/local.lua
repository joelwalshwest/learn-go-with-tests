-- Enable go testing. Workaround for multiple test adapters not working
require("neotest").setup({adapters = {require("neotest-golang")}})
