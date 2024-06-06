

    
export interface EndPoint {
    name: string;

}

export interface Service {
    name: string;

    endpoints: EndPoint[]
}

export interface AppMap {
    name: string;

    services: Service[]
}
export const pagemap : AppMap = {
  "services": [
    {
      "endpoints": [
        {
          "name": "read"
        },
        {
          "name": "create"
        },
        {
          "name": "update"
        },
        {
          "name": "delete"
        },
        {
          "name": "search"
        }
      ],
      "name": "category"
    },
    {
      "endpoints": [
        {
          "name": "read"
        },
        {
          "name": "create"
        },
        {
          "name": "update"
        },
        {
          "name": "delete"
        },
        {
          "name": "search"
        }
      ],
      "name": "tool"
    },
    {
      "endpoints": [
        {
          "name": "read"
        },
        {
          "name": "create"
        },
        {
          "name": "update"
        },
        {
          "name": "delete"
        },
        {
          "name": "search"
        }
      ],
      "name": "region"
    },
    {
      "endpoints": [
        {
          "name": "read"
        },
        {
          "name": "create"
        },
        {
          "name": "update"
        },
        {
          "name": "delete"
        },
        {
          "name": "search"
        }
      ],
      "name": "country"
    },
    {
      "endpoints": [
        {
          "name": "read"
        },
        {
          "name": "create"
        },
        {
          "name": "update"
        },
        {
          "name": "delete"
        },
        {
          "name": "search"
        }
      ],
      "name": "user"
    },
    {
      "endpoints": [
        {
          "name": "read"
        },
        {
          "name": "create"
        },
        {
          "name": "update"
        },
        {
          "name": "delete"
        },
        {
          "name": "search"
        }
      ],
      "name": "language"
    },
    {
      "endpoints": [
        {
          "name": "read"
        },
        {
          "name": "create"
        },
        {
          "name": "update"
        },
        {
          "name": "delete"
        },
        {
          "name": "search"
        }
      ],
      "name": "translation"
    },
    {
      "endpoints": [
        {
          "name": "read"
        },
        {
          "name": "create"
        },
        {
          "name": "update"
        },
        {
          "name": "delete"
        },
        {
          "name": "search"
        }
      ],
      "name": "icon"
    },
    {
      "endpoints": [
        {
          "name": "read"
        },
        {
          "name": "create"
        },
        {
          "name": "update"
        },
        {
          "name": "delete"
        },
        {
          "name": "search"
        }
      ],
      "name": "attachment"
    },
    {
      "endpoints": [
        {
          "name": "read"
        },
        {
          "name": "create"
        },
        {
          "name": "update"
        },
        {
          "name": "delete"
        },
        {
          "name": "search"
        }
      ],
      "name": "region"
    }
  ],
  "name": "people"
}

