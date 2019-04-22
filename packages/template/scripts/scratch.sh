#!/bin/bash

dao apps 0xE53Cedef126e92233Aba07655318A4C0D27BCdb6 | awk 'NR>1 { if ($2 == "acl") print $4 }'
