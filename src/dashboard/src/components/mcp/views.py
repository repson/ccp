# This file is part of Archivematica.
#
# Copyright 2010-2013 Artefactual Systems Inc. <http://artefactual.com>
#
# Archivematica is free software: you can redistribute it and/or modify
# it under the terms of the GNU Affero General Public License as published by
# the Free Software Foundation, either version 3 of the License, or
# (at your option) any later version.
#
# Archivematica is distributed in the hope that it will be useful,
# but WITHOUT ANY WARRANTY; without even the implied warranty of
# MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
# GNU General Public License for more details.
#
# You should have received a copy of the GNU General Public License
# along with Archivematica.  If not, see <http://www.gnu.org/licenses/>.
from client import get_client
from django.http import HttpResponse
from lxml import etree


def execute(request):
    result = ""
    if request.POST.get("uuid"):
        client = get_client(request.user.id)
        result = client.approve_job(
            request.POST.get("uuid"), request.POST.get("choice", "")
        )
    return HttpResponse(result, content_type="text/plain")


def list(request):
    client = get_client(request.user.id)
    jobs = etree.XML(client.list_jobs_awaiting_approval())
    response = ""
    if 0 < len(jobs):
        for job in jobs:
            response += etree.tostring(job, encoding="utf8").decode("utf8")
    response = "<MCP>%s</MCP>" % response
    return HttpResponse(response, content_type="text/xml")
