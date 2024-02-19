
# ApiDebugger
This project is exclusively a way for me to try new things.
It is supposed to become a basic tool for debugging APIs with a simple web-ui.
Technology used: Golang, HTMX, TailwindCSS, DaisyUI, YAML.

The following format is used to configure the different API endpoints you want to debug:
```YAML
headers:
  Content-Type:
    # the contents of this section describe an "InputField"
    type: dropdown
    description: headers in the root section apply to every single request
    locked: false
    options:
      - application/json
      - text/html
    default: application/json
services:
  some-service:
    environments:
      dev:
        url: localhost:8080
        paths:
          /hello/{name}:
            description: personalized greeting handler
            variables:
              name:
                # InputField
            methods:
              GET:
                # InputField
```
